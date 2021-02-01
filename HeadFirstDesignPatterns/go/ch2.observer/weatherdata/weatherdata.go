package weatherdata

type Observer interface {
	Update(temp float32, humidity float32, pressure float32)
}

type Subject interface {
	RegisterObserver(o Observer)
	removeObserver(o *Observer)
	notifyObserver()
}

type WeatherData struct {
	Temperature float32
	Humidity    float32
	Pressure    float32
	observers   []*Observer
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, &o)
	w.Temperature = 0
}

func (w *WeatherData) removeObserver(o *Observer) {
	var n []*Observer
	for _, i := range w.observers {
		if i != o {
			n = append(n, i)
		}
	}
	w.observers = n
}

func (w *WeatherData) notifyObserver() {
	var o *Observer
	for _, o = range w.observers {
		(*o).Update(w.Temperature, w.Humidity, w.Pressure)
	}
}

func (w *WeatherData) measurementsChanged() {
	w.notifyObserver()
}

func (w *WeatherData) SetMeasurements(temperature float32, humidity float32, pressure float32) {
	w.Temperature = temperature
	w.Humidity = humidity
	w.Pressure = pressure

	w.measurementsChanged()
}
