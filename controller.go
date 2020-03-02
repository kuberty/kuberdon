package main
const controllerAgentName = "kuberdon-controller"

const (
	StateSynced = "Synced"
	StateSyncing = "Syncing"
        StateError = "Error"

	MessageSecretAlreadyExists = "Secret %q already exists and is not managed by this kuberdon controller."
	MessageSynced = "Kuberdon synced %d	secrets."
)



