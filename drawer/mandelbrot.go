package drawer

type MandelbrotGenerator interface {
	Generate(canvas *Image) error
	GenerateParallel(canvas *Image) error
}

func AutoFitScaleFactor(res int) float32 {
	return 3.5 / float32(res)
}
