import sys
import csv
import os
import validacion as v
import ejecuciones as e
ARCHIVO = "reservas.csv"

def main():
    validado, mensaje, comando = v.reconocer_mensaje(ARCHIVO)
    if validado:
        e.realizar_comando(mensaje, comando, ARCHIVO)
    
if __name__ == "__main__":
    main()