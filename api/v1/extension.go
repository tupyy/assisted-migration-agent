package v1

import (
	"github.com/tupyy/assisted-migration-agent/internal/models"
)

func (a *AgentStatus) FromModel(m models.AgentStatus) {
	a.ConsoleConnection = AgentStatusConsoleConnection(m.Console.Current)
	a.Mode = AgentStatusMode(m.Console.Target)
}
