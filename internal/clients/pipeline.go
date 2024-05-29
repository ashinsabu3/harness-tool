package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ashinsabu/harness-tool/internal/config_parser"
	"gopkg.in/yaml.v3"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Data   struct {
		PlanExecution struct {
			UUID string `json:"uuid"`
		} `json:"planExecution"`
	} `json:"data"`
}

// ExecutePipeline sends a request to the Harness API to execute a pipeline
func ExecutePipeline(harnessConfig *config_parser.HarnessConfig, pipelineID, inputSetID string) (string, error) {
	// Construct the API URL with query parameters
	apiURL := fmt.Sprintf(
		"https://app.harness.io/gateway/pipeline/api/pipeline/execute/%s?routingId=%s&accountIdentifier=%s&projectIdentifier=%s&orgIdentifier=%s",
		pipelineID,
		harnessConfig.HarnessMeta.AccountId,
		harnessConfig.HarnessMeta.AccountId,
		harnessConfig.HarnessMeta.ProjectId,
		harnessConfig.HarnessMeta.OrgId,
	)

	// Prepare the pipeline data
	pipelineData := map[string]interface{}{
		"pipeline": map[string]interface{}{
			"identifier": pipelineID,
			"variables":  buildVariables(harnessConfig, pipelineID, inputSetID),
			"properties": harnessConfig.HarnessProperties,
		},
	}

	// Convert the pipeline data to YAML
	pipelineYAML, err := yaml.Marshal(pipelineData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal pipeline data to YAML: %w", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(pipelineYAML))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/yaml")
	req.Header.Set("x-api-key", harnessConfig.HarnessMeta.ApiKey)

	// Create an HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK response: %s", resp.Status)
	}

	// Decode the JSON response
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to decode JSON response: %w", err)
	}

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK response: %s", resp.Status)
	}

	return response.Data.PlanExecution.UUID, nil
}

// buildVariables constructs the variables section from the HarnessConfig struct
func buildVariables(harnessConfig *config_parser.HarnessConfig, pipelineID, inputSetID string) []map[string]interface{} {
	var variables []map[string]interface{}
	for _, pipeline := range harnessConfig.HarnessPipelines {
		if pipeline.Id == pipelineID {
			for _, inputSet := range pipeline.LocalInputSets {
				if inputSet.Id == inputSetID {
					for _, variable := range inputSet.InputsetVars {
						variables = append(variables, map[string]interface{}{
							"name":  variable.Name,
							"type":  variable.Type,
							"value": variable.Value,
						})
					}
				}
			}
		}
	}
	return variables
}
