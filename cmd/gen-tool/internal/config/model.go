/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-21 11:36:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-31 18:13:52
 * @Description: model config
 */
package config

import (
	"fmt"
	"strings"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/errors"
)

// Model struct represents the configuration for a model.
// It holds the CLI context and the model name.
// The CLI context is used to access command-line arguments and options.
type Model struct {
	// CliContext holds the CLI context which contains information about the command-line invocation.
	CliContext *cli.Context
	// TableName
	TableName string
	// ModelName stores the name of the model.
	ModelName string
}

/**
 * @description: ModelConf
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:37:21
 * @return {*}
 */
// ModelConf is a global variable representing the model configuration.
// It is initialized as a pointer to a Model struct with default values.
var ModelConf = &Model{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:38:59
 * @return {*}
 */
// SetConfig is a method of the Model struct.
// It is responsible for setting the model configuration based on the command-line arguments.
// If successful, it updates the ModelConf.ModelName with the processed model name.
// Returns an error if any step in the configuration setting process fails.
func (m *Model) SetConfig() error {

	// Check if the CliContext is nil.
	// If it is, return an error indicating that the CLI context is missing.
	if m.CliContext == nil {
		return errors.ErrConfigNoCliContext
	}

	// Get the first argument from the command-line arguments.
	// This first argument is considered as the model name.
	modelName := m.CliContext.Args().First()

	// Check if the model name is empty.
	// If it is, return an error indicating that the model name is not provided.
	if modelName == "" {
		return errors.ErrConfigNoModelName
	}

	// Set TableName
	ModelConf.TableName = modelName

	// Split the model name by underscore into an array of words.
	modelNameArr := strings.Split(modelName, "_")
	// Check if the resulting array is empty.
	// If it is, return an error indicating that the model name is not in a valid format.
	if len(modelNameArr) == 0 {
		return errors.ErrConfigNoModelName
	}

	// Initialize the model name string with the first word from the array.
	modelNameStr := modelNameArr[0]
	// Iterate over the remaining words in the array.
	for i := 1; i < len(modelNameArr); i++ {
		// Check if the word has a length greater than 0.
		// If so, convert the first letter of the word to uppercase and append it to the model name string.
		if len(modelNameArr[i]) > 0 {
			modelNameStr += strings.Title(modelNameArr[i])
		}
	}

	// Set the model name in the ModelConf with the processed model name prefixed with "Model".
	ModelConf.ModelName = fmt.Sprintf("%sModel", modelNameStr)
	return nil
}
