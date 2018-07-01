package rabc

type Feature struct {
}

func (t *Feature) Has(userFeatue int, f1 int) bool {
	return userFeatue&f1 > 0
}
func (t *Feature) Grant(userFeatue int, f1 int) int {
	return userFeatue | f1
}
func (t *Feature) Revoke(userFeatue int, f1 int) int {
	return userFeatue ^ f1
}
func NewFeature() *Feature {
	return &Feature{}
}
