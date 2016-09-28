package math
func Averange(xs []float64) float64{
  total:=0.0
  for _, x:=range xs{
    total+=x
  }
  return total/float64(len(xs))
}
