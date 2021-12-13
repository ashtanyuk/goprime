package goprime

func isPrime(number int) bool {
   if n<=1 {
     return false
   }
   if n == 2 {
     return true
   }
   for i := 2; i <= number; i++ {
     if number % i == 0 {
        return false
     }
   }
   return true
}