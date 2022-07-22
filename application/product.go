package application

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ANABLED  = "anabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

func (p *Product) IsValid() (bool, error) {

}

func (p *Product) Enable() error {

}

func (p *Product) Disable() error {

}

func (p *Product) GetID() string {

}

func (p *Product) GetName() string {

}

func (p *Product) GetPrice() float64 {

}

func (p *Product) GetStatus() string {

}