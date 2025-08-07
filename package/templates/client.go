package {{ .Values.package }}


type Client struct {
  // Add any fields that are necessary for the client
}

// New constructor function for creating a new instance of Client
func New() *Client {
  return &Client{
    // Initialize any fields if necessary
  }
}

