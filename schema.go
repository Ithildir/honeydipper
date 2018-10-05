package main

import (
	"time"
)

// Event : the runtime data representation of an event
type Event struct {
	System  string
	Trigger string
}

// Action : the runtime data representation of an action
type Action struct {
	System   string
	Function string
}

// Filter : internal operation to mutate the payload data between events and actions
type Filter interface{}

// Trigger : is the datastructure hold the information to match and process an event
type Trigger struct {
	Driver     string                   `yaml:"driver,omitempty"`
	RawEvent   string                   `yaml:"rawevent,omitempty"`
	Conditions (interface{})            `yaml:"conditions,omitempty"`
	Fields     map[string](interface{}) `yaml:"fields,omitempty"`
	Base       Event                    `yaml:"base,omitempty"`
	Filters    []Filter                 `yaml:"filters,omitempty"`
}

// Function : is the datastructure hold the information to run actions
type Function struct {
	Driver     string                   `yaml:"driver,omitempty"`
	RawAction  string                   `yaml:"rawaction,omitempty"`
	Parameters map[string](interface{}) `yaml:"parameters,omitempty"`
	Results    map[string](interface{}) `yaml:"results,omitempty"`
	Base       Action                   `yaml:"base,omitempty"`
	Filters    []Filter                 `yaml:"filters,omitempty"`
}

// System : is an abstract construct to group data, trigger and function definitions
type System struct {
	Data      (interface{})       `yaml:"data,omitempty"`
	Triggers  map[string]Trigger  `yaml:"triggers,omitempty"`
	Functions map[string]Function `yaml:"functions,omitempty"`
}

// Condition : used to for conditioning in workflow
type Condition struct {
	Op     string `yaml:"op,omitempty"`
	Values []string
}

// Workflow : defines the steps, and relationship of the actions
type Workflow struct {
	Block      string
	Conditions []Condition `yaml:"conditions,omitempty"`
	Content    [](interface{})
}

// Rule : is a data structure defining what action to take upon an event
type Rule struct {
	When Trigger
	Do   Function
}

// DriverData : holds the data necessary for a driver to operate
type DriverData interface{}

// RepoInfo : points a git repo where config data can be read from
type RepoInfo struct {
	Repo   string
	Branch string `yaml:"branch,omitempty"`
	Path   string `yaml:"path,omitempty"`
}

// ConfigSet : is a complete set of configuration at a specific moment
type ConfigSet struct {
	Systems  map[string]System     `yaml:"systems,omitempty"`
	Rules    map[string]Rule       `yaml:"rules,omitempty"`
	Drivers  map[string]DriverData `yaml:"drivers,omitempty"`
	Includes []string              `yaml:"includes,omitempty"`
	Repos    []RepoInfo            `yaml:"repos,omitempty"`
}

// ConfigRepoRuntime : used to track what has been loaded in a repo
type ConfigRepoRuntime struct {
	parent *ConfigRuntime
	config ConfigSet
	files  map[string]bool
	root   string
}

// ConfigRuntime : used to track what has been loaded
type ConfigRuntime struct {
	config ConfigSet
	loaded map[RepoInfo]ConfigRepoRuntime
	wd     string
}

// Config : is the complete configration of the daemon including history and the running services
type Config struct {
	initRepo RepoInfo
	services []string
	revs     map[time.Time]ConfigRuntime
}