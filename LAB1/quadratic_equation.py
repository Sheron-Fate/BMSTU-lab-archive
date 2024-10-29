import math


class Equation:
    def __init__(self, abc: list):
        self.A: int = abc[0]
        self.B: int = abc[1]
        self.C: int = abc[2]
   
    def __str__(self):
        return f"{self.A}x^4 + {self.B}x^2  + {self.C}"

    def get_discriminant(self) -> int:
        return self.B**2 - 4*self.A*self.C

    def get_quadratic_roots(self) -> list:
        if self.get_discriminant() == 0:
            return [(-self.B) / 2*self.A]
        elif self.get_discriminant() > 0:
            return [(-self.B + math.sqrt(self.get_discriminant())) / 2*self.A,
                    (-self.B - math.sqrt(self.get_discriminant())) / 2*self.A]
        else:
            return []
        
    def get_biquadratic_roots(self, roots: list) -> list:
        broots = []
        for i in range(len(roots)):
            if roots[i] > 0:
                broots.append(roots[i]**0.5)
                broots.append((-1)*(roots[i]**0.5))
        return broots
