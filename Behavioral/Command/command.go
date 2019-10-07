// Command пакет є прикладом шаблону команд.
package Command

// Command інтерфейс команд.
type Command interface {
	Execute() string
}

// Receiver реалізація.
type Receiver struct {}

// ToggleOn
func (*Receiver) ToggleOn() string {
	return "Toggle on!"
}

// ToggleOff
func (*Receiver) ToggleOff() string {
	return "Toggle off!"
}

// ToggleOnCommand реалізує інтерфейс Command
type ToggleOnCommand struct {
	receiver *Receiver
}

// Execute виконати команду
func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

// ToggleOffCommand реалізує інтерфейс Command
type ToggleOffCommand struct {
	receiver *Receiver
}

// Execute виконати команду
func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

// Invoker
type Invoker struct {
	commands []Command
}

// StoreCommand додати команди до стеку
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// UnStoreCommand видалити команди зі стеку
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Execute виконати всі команди
func (i *Invoker) Execute() string {
	var result string
	for _, v :=  range i.commands {
		result += v.Execute() + "\n"
	}
	return result
}