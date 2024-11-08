package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Systenix/go-cloud/templates"
)

type Field struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	JSONName string `yaml:"json_name"`
}

type Model struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`
}

type Repository struct {
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

type Route struct {
	Path          string   `yaml:"path"`
	Middleware    []string `yaml:"middleware"`
	Method        string   `yaml:"method"`
	Function      string   `yaml:"function"`
	ServiceMethod string   `yaml:"service_method"`
	RequestModel  string   `yaml:"request_model"`
	ResponseModel string   `yaml:"response_model"`
}

type Handler struct {
	Name    string  `yaml:"name"`
	Service string  `yaml:"service"`
	Routes  []Route `yaml:"routes"`
}

type Param struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type ServiceMethod struct {
	Name    string  `yaml:"name"`
	Params  []Param `yaml:"params"`
	Returns []Param `yaml:"returns"`
}

type Service struct {
	Name         string          `yaml:"name"`
	Type         string          `yaml:"type"`
	Port         string          `yaml:"port"`
	Models       []Model         `yaml:"models"`
	ModelNames   []string        `yaml:"model_names"`
	Repositories []Repository    `yaml:"repositories"`
	Handlers     []Handler       `yaml:"handlers"`
	Middleware   []string        `yaml:"middleware"`
	Methods      []ServiceMethod `yaml:"methods"`
}

type Middleware struct {
	Name    string            `yaml:"name"`
	Type    string            `yaml:"type"`
	Scope   string            `yaml:"scope"`
	Options map[string]string `yaml:"options"`
}

type Prometheus struct {
	Enabled bool   `yaml:"enabled"`
	Port    string `yaml:"port"`
}

type Grafana struct {
	Enabled bool   `yaml:"enabled"`
	Port    string `yaml:"port"`
}

type ThirdParty struct {
	Prometheus Prometheus `yaml:"prometheus"`
	Grafana    Grafana    `yaml:"grafana"`
}

type Docker struct {
	Enabled bool `yaml:"enabled"`
}

type Event struct {
	Name        string `yaml:"name"`
	Payload     string `yaml:"payload"`
	Description string `yaml:"description"`
}

type ProjectData struct {
	ProjectName   string
	ProjectPath   string
	ModulePath    string
	ProjectDir    string
	Protocol      string
	MessageBroker string
	GoVersion     string
	Middleware    []Middleware
	ThirdParty    ThirdParty
	Services      []Service
	Models        []Model
	Events        []Event
	Docker        Docker
}

func GenerateProject(data ProjectData) error {
	data.ModulePath = filepath.Join(data.ProjectPath, data.ProjectName)
	data.ProjectDir = filepath.Join("./generated", data.ModulePath)

	// Create project directory
	err := os.MkdirAll(data.ProjectDir, 0755)
	if err != nil {
		return err
	}

	// Generate go.mod
	err = generateFile("go.mod.tmpl", filepath.Join(data.ProjectDir, "go.mod"), data)
	if err != nil {
		return err
	}

	// Generate middleware
	err = generateMiddleware(data)
	if err != nil {
		return err
	}

	// Generate services
	err = generateServices(data)
	if err != nil {
		return err
	}

	// Generate handlers
	err = generateHandlers(data)
	if err != nil {
		return err
	}

	// Generate models
	err = generateModels(data)
	if err != nil {
		return err
	}

	// Generate events
	err = generateEvents(data)
	if err != nil {
		return err
	}

	// Generate main.go
	err = generateMain(data)
	if err != nil {
		return err
	}

	// Generate Dockerfile if enabled
	if data.Docker.Enabled {
		err = generateFile("Dockerfile.tmpl", filepath.Join(data.ProjectDir, "Dockerfile"), data)
		if err != nil {
			return err
		}
	}

	// Generate docker-compose.yml if enabled
	if data.Docker.Enabled {
		err = generateFile("docker-compose.yml.tmpl", filepath.Join(data.ProjectDir, "docker-compose.yml"), data)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateMain(data ProjectData) error {
	// Collect imports
	imports := []string{
		"log",
		"github.com/gin-gonic/gin",
		fmt.Sprintf("%s/internal/interfaces/handlers", data.ModulePath),
		fmt.Sprintf("%s/internal/interfaces/middleware", data.ModulePath),
		fmt.Sprintf("%s/internal/services", data.ModulePath),
	}

	middleware := collectMiddleware(data)

	fmt.Println("main Middleware:", middleware)

	// Prepare template data
	tmplData := struct {
		ModulePath string
		Services   []Service
		Handlers   []Handler
		Middleware []Middleware
		Imports    []string
	}{
		ModulePath: data.ModulePath,
		Services:   data.Services,
		Handlers:   collectHandlers(data.Services),
		Middleware: middleware,
		Imports:    imports,
	}

	outputPath := filepath.Join(data.ProjectDir, "cmd", "main.go")
	return generateFile("cmd/main.go.tmpl", outputPath, tmplData)
}

func collectMiddleware(data ProjectData) []Middleware {
	var middlewareList []Middleware
	middlewareMap := make(map[string]Middleware)

	// Collect middleware from global definitions
	for _, mw := range data.Middleware {
		middlewareMap[mw.Name] = mw
		middlewareList = appendIfMissing(middlewareList, mw)
	}

	// Collect middleware from services, handlers, and routes
	for _, service := range data.Services {
		// Service-level middleware (if any)
		for _, mwName := range service.Middleware {
			if mw, exists := middlewareMap[mwName]; exists {
				middlewareList = appendIfMissing(middlewareList, mw)
			}
		}
		for _, handler := range service.Handlers {
			for _, route := range handler.Routes {
				// Route-level middleware (if any)
				for _, mwName := range route.Middleware {
					if mw, exists := middlewareMap[mwName]; exists {
						middlewareList = appendIfMissing(middlewareList, mw)
					}
				}
			}
		}
	}

	return middlewareList
}

func appendIfMissing(middlewareList []Middleware, mw Middleware) []Middleware {
	for _, existingMw := range middlewareList {
		if existingMw.Name == mw.Name {
			return middlewareList
		}
	}
	return append(middlewareList, mw)
}

func collectRepositories(services []Service) []Repository {
	var repositories []Repository
	for _, service := range services {
		repositories = append(repositories, service.Repositories...)
	}
	return repositories
}

// Collect handlers from all services
func collectHandlers(services []Service) []Handler {
	var handlers []Handler
	for _, service := range services {
		handlers = append(handlers, service.Handlers...)
	}
	return handlers
}

func generateMiddleware(data ProjectData) error {
	if len(data.Middleware) == 0 {
		// No middleware to generate
		return nil
	}

	// Determine necessary imports based on middleware types
	importsSet := make(map[string]struct{})
	for _, mw := range data.Middleware {
		switch mw.Type {
		case "logging":
			importsSet["log"] = struct{}{}
			importsSet["time"] = struct{}{}
		case "metrics":
			importsSet["github.com/prometheus/client_golang/prometheus"] = struct{}{}
			importsSet["github.com/prometheus/client_golang/prometheus/promhttp"] = struct{}{}
			// Add other types as necessary
		}
	}

	// Convert imports set to slice
	imports := make([]string, 0, len(importsSet))
	for imp := range importsSet {
		imports = append(imports, imp)
	}

	tmplData := struct {
		Middleware []Middleware
		Imports    []string
	}{
		Middleware: data.Middleware,
		Imports:    imports,
	}

	outputPath := filepath.Join(data.ProjectDir, "internal", "interfaces", "middleware", "middleware.go")
	return generateFile("interfaces/middleware/middleware.go.tmpl", outputPath, tmplData)
}

// Generate services for a project
func generateServices(data ProjectData) error {
	for i := range data.Services {
		service := &data.Services[i]

		// Generate models for the service
		err := generateServiceModels(data, service)
		if err != nil {
			return err
		}

		// Generate repositories for the service
		err = generateServiceRepositories(data, service)
		if err != nil {
			return err
		}

		// Determine required imports
		imports := getServiceImports(service)

		// Always include mandatory imports
		mandatoryImports := []string{
			fmt.Sprintf("%s/internal/repositories", data.ModulePath),
			fmt.Sprintf("%s/internal/models", data.ModulePath),
		}

		// Combine mandatory and optional imports
		allImports := append(mandatoryImports, imports...)

		// Generate service code
		outputPath := filepath.Join(data.ProjectDir, "internal", "services", service.Name+".go")
		tmplData := struct {
			ModulePath string
			Service    Service
			Imports    []string
		}{
			ModulePath: data.ModulePath,
			Service:    *service,
			Imports:    allImports,
		}
		err = generateFile("services/service.go.tmpl", outputPath, tmplData)
		if err != nil {
			return err
		}
	}
	return nil
}

// Generate models for a service, leveraging  shared or predefined models
func generateServiceModels(data ProjectData, service *Service) error {
	modelMap := make(map[string]Model)
	for _, model := range data.Models {
		modelMap[model.Name] = model
	}

	for _, modelName := range service.ModelNames {
		model, exists := modelMap[modelName]
		if !exists {
			return fmt.Errorf("model '%s' not found in project models", modelName)
		}

		imports := getModelImports(model)

		tmplData := struct {
			ModulePath string
			Model      Model
			Imports    []string
		}{
			ModulePath: data.ModulePath,
			Model:      model,
			Imports:    imports,
		}

		outputPath := filepath.Join(data.ProjectDir, "internal", "models", model.Name+".go")
		err := generateFile("models/model.go.tmpl", outputPath, tmplData)
		if err != nil {
			return err
		}
	}
	return nil
}

// Generate repositories for a service
func generateServiceRepositories(data ProjectData, service *Service) error {
	for _, repo := range service.Repositories {
		tmplData := struct {
			ModulePath string
			Repository Repository
		}{
			ModulePath: data.ModulePath,
			Repository: repo,
		}

		outputPath := filepath.Join(data.ProjectDir, "internal", "infrastructures", "repositories", repo.Name+".go")
		err := generateFile("infrastructures/repositories/repository.go.tmpl", outputPath, tmplData)
		if err != nil {
			return err
		}
	}
	return nil
}

// Determine imports needed by a service
func getServiceImports(service *Service) []string {
	importsSet := make(map[string]struct{})

	for _, method := range service.Methods {
		for _, param := range method.Params {
			switch param.Type {
			case "context.Context":
				importsSet["context"] = struct{}{}
				// Add other standard library types as needed
			}
		}
		for _, ret := range method.Returns {
			switch ret.Type {
			case "error":
				// No import needed for error
				// Add other standard library types as needed
			}
		}
	}

	imports := make([]string, 0, len(importsSet))
	for imp := range importsSet {
		imports = append(imports, imp)
	}

	return imports
}

// Generate handlers for a service
func generateHandlers(data ProjectData) error {
	for _, service := range data.Services {
		for _, handler := range service.Handlers {
			// Determine imports needed by the handler
			imports := getHandlerImports(handler, data.ModulePath)

			outputPath := filepath.Join(data.ProjectDir, "internal", "interfaces", "handlers", handler.Name+".go")
			tmplData := struct {
				ModulePath string
				Handler    Handler
				Service    Service
				Imports    []string
			}{
				ModulePath: data.ModulePath,
				Handler:    handler,
				Service:    service,
				Imports:    imports,
			}
			err := generateFile("interfaces/handlers/handler.go.tmpl", outputPath, tmplData)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Determine imports needed by a handler
func getHandlerImports(handler Handler, modulePath string) []string {
	importsSet := make(map[string]struct{})

	// Always include necessary imports
	importsSet["net/http"] = struct{}{}
	importsSet["github.com/gin-gonic/gin"] = struct{}{}
	importsSet[fmt.Sprintf("%s/internal/services", modulePath)] = struct{}{}

	// Include models if needed
	for _, route := range handler.Routes {
		if route.RequestModel != "" || route.ResponseModel != "" {
			importsSet[fmt.Sprintf("%s/internal/models", modulePath)] = struct{}{}
			break
		}
	}

	// Convert set to slice
	imports := make([]string, 0, len(importsSet))
	for imp := range importsSet {
		imports = append(imports, imp)
	}

	return imports
}

// Determine imports needed by a model
func getModelImports(model Model) []string {
	importsSet := make(map[string]struct{})

	for _, field := range model.Fields {
		switch field.Type {
		case "time.Time":
			importsSet["time"] = struct{}{}
			// Add more types as needed
		}
	}

	imports := make([]string, 0, len(importsSet))
	for imp := range importsSet {
		imports = append(imports, imp)
	}

	return imports
}

// Generate models for a project
func generateModels(data ProjectData) error {
	for _, model := range data.Models {
		imports := getModelImports(model)

		tmplData := struct {
			ModulePath string
			Model      Model
			Imports    []string
		}{
			ModulePath: data.ModulePath,
			Model:      model,
			Imports:    imports,
		}

		outputPath := filepath.Join(data.ProjectDir, "internal", "models", model.Name+".go")
		err := generateFile("models/model.go.tmpl", outputPath, tmplData)
		if err != nil {
			return err
		}
	}
	return nil
}

// Generate events for a project
func generateEvents(data ProjectData) error {
	for _, event := range data.Events {
		tmplData := struct {
			ModulePath string
			Event      Event
		}{
			ModulePath: data.ModulePath,
			Event:      event,
		}

		outputPath := filepath.Join(data.ProjectDir, "internal", "events", event.Name+".go")
		err := generateFile("events/event.go.tmpl", outputPath, tmplData)
		if err != nil {
			return err
		}
	}
	return nil
}

// Generate a file from a template
func generateFile(templatePath, outputPath string, data interface{}) error {
	fmt.Println("Generating file", outputPath)
	// fmt.Println("Passed data", data)

	tmplContent, err := templates.FS.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", templatePath, err)
	}

	// Create a new template and add the custom functions
	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(template.FuncMap{
		"lowerFirst": lowerFirst,
	}).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	// Ensure the output directory exists
	err = os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", filepath.Dir(outputPath), err)
	}

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer outFile.Close()

	// Execute the template
	return tmpl.Execute(outFile, data)
}

// lowerFirst converts the first character of a string to lowercase.
func lowerFirst(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToLower(str[:1]) + str[1:]
}
