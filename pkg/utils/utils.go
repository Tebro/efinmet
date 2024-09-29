package utils

import (
	"math"
	"os"
	"os/exec"
	"runtime"
	"strings"
)


func FieldIcaoHasPrefix(prefix string, fieldIcao string) bool {
	return strings.HasPrefix(fieldIcao, prefix)
}


func ClearTerm() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear")
		break
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
		break
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DistanceBetween(lat1, lon1, lat2, lon2 float64) float64 {
	r := 6371.0 // Radius of the Earth in km
	x := (lon2 - lon1) * math.Pi / 180 * math.Cos(((lat1+lat2)/2)*math.Pi/180)
	y := (lat2 - lat1) * math.Pi / 180
	return r * math.Sqrt(x*x+y*y)
}
