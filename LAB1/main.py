from quadratic_equation import Equation
import sys


def get_ABC() -> list:
    abc = []
    for i in range(1, 4):
        try:
            arg = sys.argv[i]
        except:
            arg = int(input('Enter an integer: '))

        abc.append(arg)
    return abc


if __name__ == '__main__':

    abc = get_ABC()
    equation = Equation(abc)
    roots = equation.get_biquadratic_roots(equation.get_quadratic_roots())
    print(roots)
