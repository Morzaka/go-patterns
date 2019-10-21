// State пакет реалізує приклад патерну Стану
package State

// MobileAlertStater забезпечує загальний інтерфейс для різних станів.
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert реалізує Alert залежно від свого стану.
type MobileAlert struct {
	state MobileAlertStater
}

// Alert повертає рядок попередження
func (m *MobileAlert) Alert() string {
	return m.state.Alert()
}

// SetState змінює стан
func (m *MobileAlert) SetState(state MobileAlertStater) {
	m.state = state
}

// NewMobileAlert - це MobileAlert конструктор
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration реалізує вібро сигнал
type MobileAlertVibration struct {}
func (m *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// MobileAlertVibration реалізує сигнал оповіщення
type MobileAlertSong struct {}
func (m *MobileAlertSong) Alert() string {
	return "Черешні рвати, черешні їсти..."
}