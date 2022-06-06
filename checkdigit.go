package checkdigit


type (
    Verifier interface {
        Verify(code string) bool
    }
    Generator interface {
        Generate(seed string) (int, error)
    }
    Provider interface {
        Verifier
        Generator
    }
)

func NewLuhn() Provider {
    return &luhn{}
}

func isNotNumber(n rune) bool {
    return n < '0' || '9' < n
}