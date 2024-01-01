import sys
import csv
import os

CAMBIO_PEDIDO = 1
CAMPO_PEDIDO = 0
#FORMATO
CAMPOS = ["id","nombre","cant","hora","ubicacion"]
ID=0
NOMBRE = 1
CANT_PERSONAS = 2
HORARIO = 3
UBICACION = 4

#TIPOS DE FUNCIONES
FUNCION_AGREGRAR = "agregar"
FUNCION_MODIFICAR = "modificar"
FUNCION_ELIMINAR = "eliminar"
FUNCION_LISTAR = "listar"

#CANTIDADES
CANTIDAD_ARGUMENTOS_CAMBIO = 2
CANTIDAD_MINIMA_DE_PERSONAS = 1
CANTIDAD_ARGUMENTOS_AGREGAR = 6
CANTIDAD_ARGUMENTOS_MODIFICAR_ELIMINAR = 3
CANTIDAD_DE_ARGUMENTOS_LISTAR_ENTERO = 2
CANTIDAD_DE_ARGUMENTOS_LISTAR_PARAMETRIZADO = 4
CANTIDAD_MINIMO_RANGO = 0

#HORARIO
LARGO_HORARIO = 5
MITAD_HORARIO = 2
MEDIO_HORA = ":"
MINIMO_HORARIO = 0
MAXIMO_HORAS = 23
MAXIMO_MINUTOS = 59

#UBICACIONES
AFUERA = "F"
ADENTRO = "D"

#POSICIONES PARA AGREGAR
POSICION_FUNCION = 1
POSICION_NOMBRE = 2
POSICION_CANTIDAD_PERSONAS = 3
POSICION_HORARIO = 4
POSICION_UBICACION = 5

#POSICIONES PARA LISTAR
RANGO_MIN_INDICADO = 2
RANGO_MAX_INDICADO = 3


#-HERRAMIENTAS COMUNES-

#PRE: 
#POST: Verifica que el .csv dado exista, si el comando es agregar, lo creara vacio.
def verificar_archivo_existe( archivo ,comando):
        existe = True
        if not os.path.exists(archivo) and comando == FUNCION_AGREGRAR:
            archivo = open(archivo, 'w')
            archivo.close()
        elif not os.path.exists(archivo):
            print("No existen reservas, comience agregando una")
            existe = False
        return existe

#PRE: En caso de solo tener una cantidad posible, se le dara la misma dos veces
#POST: recibe una lista y verifica que su largo sea/sean el/los indicado/s
def validar_cantidad_de_argumentos(mensaje , cantidad_de_argumentos1 , cantidad_de_argumentos2):
    validado = True
    if( len(mensaje)!= cantidad_de_argumentos1 and len(mensaje)!= cantidad_de_argumentos2):
        print("Cantidad invalida de argumentos")
        validado=False
    return validado



#-VALIDAR AGREGAR-

#PRE: 
#POST: Verifica que la cantidad de personas sea un numero mayor o igual al minimo 
def validar_cantidad_de_personas( cantidad ):
    validado = True
    if cantidad.isnumeric():
        if not(int(cantidad) >= CANTIDAD_MINIMA_DE_PERSONAS):
            print("Cantidad de personas invalida, minimo debe reservar para una persona")
            validado = False
    else:
        print("Cantidad Invalida, poné un numero")
        validado = False
    return validado

#PRE: 
#POST: Verifica que el horario este indicado correctamente segun el formato HH:MM
def validar_horario( horario ):
    hora = horario[:MITAD_HORARIO]
    minutos = horario[ (MITAD_HORARIO+1):]
    medio = horario[MITAD_HORARIO]
    validado = True
    if len(horario) == LARGO_HORARIO and medio == MEDIO_HORA:
        condicion_hora = hora.isnumeric() and int(hora)>= MINIMO_HORARIO and int(hora)<= MAXIMO_HORAS
        condicion_minutos = minutos.isnumeric() and int(minutos)>= MINIMO_HORARIO and int(minutos)<= MAXIMO_MINUTOS
        if not(condicion_hora and condicion_minutos):
            print("Horario Invalido, indique el horario entre las 00:00 y las 23:59, solamente con numeros")
            validado = False
    else:
        print("Escribiste mal el horario, utiliza el formato HH:MM\n-Por ejemplo 21:30\n")
        validado= False
    return validado

#PRE: 
#POST: Verifica que la ubicacion sea una de las dos posibles
def validar_ubicacion( mensaje ):
    validado = True
    if mensaje != AFUERA and mensaje !=ADENTRO:
        print("Ubicacion Invalida, por favor ingrese una de estas dos")
        print("F para afuera")
        print("D para adentro\n")
        validado = False
    return validado

#PRE: 
#POST: Verifica que sea posible agregar la reserva según el mensaje dado
def validar_funcion_agregar(mensaje):
    validado=True
    validado=validar_cantidad_de_argumentos(mensaje , CANTIDAD_ARGUMENTOS_AGREGAR , CANTIDAD_ARGUMENTOS_AGREGAR)
    if validado:    
        validado=validar_cantidad_de_personas(mensaje[POSICION_CANTIDAD_PERSONAS])
    if validado:   
        validado=validar_horario(mensaje[POSICION_HORARIO])
    if validado:   
        validado=validar_ubicacion(mensaje[POSICION_UBICACION])
    return validado


#-VALIDAR LISTAR-

#PRE: 
#POST: Verifica que el rango sean numeros positivos y el primero menor que el segundo. 
def validar_rango(minimo , maximo):
    validado = minimo.isnumeric() and maximo.isnumeric()
    minimo, maximo = int(minimo) , int(maximo)
    son_postivos = minimo>CANTIDAD_MINIMO_RANGO and maximo>CANTIDAD_MINIMO_RANGO
    if validado and son_postivos:
        validado = maximo>=minimo
        if not validado:
            print("El maximo debe ser mayor o igual al minimo")
    else:
        print("Debe indicar un rango con números positivos")
        validado = False
    return validado

#PRE: 
#POST: Verifica que se pueda listar segun el mensaje.
def validar_funcion_listar(mensaje):
    validado = validar_cantidad_de_argumentos(mensaje, CANTIDAD_DE_ARGUMENTOS_LISTAR_ENTERO , CANTIDAD_DE_ARGUMENTOS_LISTAR_PARAMETRIZADO )
    if validado and len(mensaje) == CANTIDAD_DE_ARGUMENTOS_LISTAR_PARAMETRIZADO:
        validado = validar_rango(mensaje[RANGO_MIN_INDICADO], mensaje[RANGO_MAX_INDICADO])
    return validado



#-PRINCIPAL-

#PRE: 
#POST: Reconoce el mensaje y el comando, verifica si se puede llevar a cabo en .csv dado
#      Comunica el error posible en caso que no se pueda llevar a cabo
def reconocer_mensaje(archivo):
    mensaje=sys.argv
    comando=mensaje[POSICION_FUNCION]
    validado = verificar_archivo_existe( archivo , comando)
    if validado:
        if comando == FUNCION_AGREGRAR:
            validado=validar_funcion_agregar(mensaje)

        elif comando == FUNCION_MODIFICAR or comando == FUNCION_ELIMINAR:
            validado = validar_cantidad_de_argumentos( mensaje, CANTIDAD_ARGUMENTOS_MODIFICAR_ELIMINAR , CANTIDAD_ARGUMENTOS_MODIFICAR_ELIMINAR)

        elif comando == FUNCION_LISTAR:
            validado=validar_funcion_listar(mensaje)

        else:
            print("Comando invalido, por favor utilice \n-agregar \n-modificar \n-eliminar \n-listar")
            print("Por favor, verifique no usar mayusculas")

    return validado, mensaje , comando



#-VALIDAR MODIFICACION PEDIDA-

#PRE: 
#POST: verifica si el campo esta correctamente escrito
def validar_campo(campo):
    campo_valido = True
    if not(campo in CAMPOS):
        print("No es un campo valido, tiene que selecionar")
        print("-nombre \n-cant \n-hora \n-ubicacion")
        campo_valido=False
    return campo_valido

#PRE: 
#POST: verifica que el nuevo valor que se quiera poner sea valido
def validar_cambio(campo , cambio):
    cambio_valido = True
    if campo == CAMPOS[CANT_PERSONAS]:
        cambio_valido=validar_cantidad_de_personas(cambio)
    elif campo == CAMPOS[HORARIO]:
        cambio_valido=validar_horario(cambio)
    elif campo == CAMPOS[UBICACION]:
        cambio_valido=validar_ubicacion(cambio)

    return cambio_valido

#PRE: 
#POST: verifica que la modificacion es posible
def validar_modificacion(modificacion):
    validado=validar_cantidad_de_argumentos(modificacion, CANTIDAD_ARGUMENTOS_CAMBIO, CANTIDAD_ARGUMENTOS_CAMBIO)
    if validado:
        validado=validar_campo(modificacion[CAMPO_PEDIDO])
    if validado:
        validado=validar_cambio(modificacion[CAMPO_PEDIDO] , modificacion[CAMBIO_PEDIDO])
    return validado


