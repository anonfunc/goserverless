package serverless

// Provider definition
type Provider struct {
	Name               string                   `json:"name,omitempty"`
	Runtime            string                   `json:"runtime,omitempty"`
	Stage              string                   `json:"stage,omitempty"`
	Region             string                   `json:"region,omitempty"`
	Profile            string                   `json:"profile,omitempty"`
	MemorySize         int                      `json:"memorySize,omitempty"`
	Timeout            int                      `json:"timeout,omitempty"`
	LogRetentionInDays int                      `json:"logRetentionInDays,omitempty"`
	DeploymentBucket   *DeploymentBucket        `json:"deploymentBucket,omitempty"`
	Role               string                   `json:"role,omitempty"`
	CFNRole            string                   `json:"cfnRole,omitempty"`
	VersionFunctions   bool                     `json:"versionFunctions,omitempty"`
	Environment        map[string]interface{}   `json:"environment,omitempty"`
	EndpointType       string                   `json:"endpointType,omitempty"`
	ResourcePolicy     []map[string]interface{} `json:"resourcePolicy,omitempty"`
	StackTags          map[string]interface{}   `json:"stackTags,omitempty"`
	Tracing            bool                     `json:"tracing,omitempty"`
	APIKeys            []string                 `json:"apiKeys,omitempty"`
	IAMRoleStatements  []map[string]interface{} `json:"iamRoleStatements,omitempty"`
	StackPolicy        []map[string]interface{} `json:"stackPolicy,omitempty"`
	NotificationARNs   []string                 `json:"notificationArns,omitempty"`
	VPC                VPC                      `json:"vpc,omitempty"`
}

// DeploymentBucket Deployment bucket name. Default is generated by the framework
type DeploymentBucket struct {
	Name                 string `json:"name,omitempty"`
	ServerSideEncryption string `json:"serverSideEncryption,omitempty"`
}
