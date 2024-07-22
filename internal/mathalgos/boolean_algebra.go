package mathalgos

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"strings"

	"github.com/Knetic/govaluate"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type LogicSimplifier struct{}

func NewLogicSimplifier() *LogicSimplifier {
	return &LogicSimplifier{}
}

func (s *LogicSimplifier) ExtractVariables(exprStr string) map[string]bool {
	variables := make(map[string]bool)
	for _, char := range exprStr {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			variables[string(char)] = true
		}
	}
	return variables
}

func (s *LogicSimplifier) TransformExpression(exprStr string) string {
	replacements := map[string]string{
		"∧": "&&",
		"∨": "||",
		"⊕": "^",
		"¬": "!",
	}
	for old, new := range replacements {
		exprStr = strings.ReplaceAll(exprStr, old, new)
	}
	return exprStr
}

type TruthTableGenerator struct {
	expression string
	simplifier *LogicSimplifier
}

func NewTruthTableGenerator(expression string) *TruthTableGenerator {
	return &TruthTableGenerator{
		expression: expression,
		simplifier: NewLogicSimplifier(),
	}
}

func (t *TruthTableGenerator) GenerateTruthTable() ([][]bool, []string, error) {
	variables := t.simplifier.ExtractVariables(t.expression)
	varNames := make([]string, 0, len(variables))
	for variable := range variables {
		varNames = append(varNames, variable)
	}

	rows := generateCombinations(len(variables))
	results := make([][]bool, len(rows))

	exprStr := t.simplifier.TransformExpression(t.expression)
	expr, err := govaluate.NewEvaluableExpression(exprStr)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing expression: %v", err)
	}

	for i, row := range rows {
		parameters := make(map[string]interface{})
		for j, variable := range varNames {
			parameters[variable] = row[j]
		}
		result, err := expr.Evaluate(parameters)
		if err != nil {
			return nil, nil, fmt.Errorf("error evaluating expression: %v", err)
		}
		results[i] = append(row, result.(bool))
	}

	return results, varNames, nil
}

func (t *TruthTableGenerator) CreateTruthTableImage() ([]byte, error) {
	truthTable, varNames, err := t.GenerateTruthTable()
	if err != nil {
		return nil, err
	}

	width := (len(varNames) + 1) * 100
	height := (len(truthTable) + 1) * 50

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	drawString(img, 10, 10, "№", black)
	for i, varName := range varNames {
		drawString(img, 100*(i+1)+10, 10, varName, black)
	}
	drawString(img, 100*(len(varNames)+1)+10, 10, t.expression, black)

	for i, row := range truthTable {
		drawString(img, 10, 50*(i+1)+10, fmt.Sprintf("%d", i), black)
		for j, val := range row {
			valStr := "0"
			if val {
				valStr = "1"
			}
			drawString(img, 100*(j+1)+10, 50*(i+1)+10, valStr, black)
		}
	}

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, img)
	if err != nil {
		return nil, fmt.Errorf("error encoding image to PNG: %v", err)
	}

	return buffer.Bytes(), nil
}

func generateCombinations(n int) [][]bool {
	combinations := [][]bool{}
	for i := 0; i < (1 << n); i++ {
		combination := make([]bool, n)
		for j := 0; j < n; j++ {
			combination[j] = (i>>j)&1 == 1
		}
		combinations = append(combinations, combination)
	}
	return combinations
}

func drawString(img *image.RGBA, x, y int, label string, col color.Color) {
	point := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
