import csv
import os
import validacion as v

ID_INDICADO = 2

#ABRIR ARCHIVO
MODO_ESCRITURA = "w"
MODO_LECTURA = "r"
MODO_AGREGAR = "a"
DELIMITER = ";"

AUXILIAR = "aux.csv"

#-HERRAMIENTAS-

#PRE: se le debe pasar un archivo.csv y el modo en el se desea abrir (escribir, agregar, leer)
#POST: trata de abrir el archivo dado en el modo pedido, si puede devuelve true
#      si no puede, lo informa y devuelve false 
def abrir_archivo( archivo , modo):
    try:
        archivo_abierto = open(archivo , modo)
        se_abrio = True
    except:
        print("No se pudo abrir el archivo")
        se_abrio = False
        archivo_abierto = None

    return archivo_abierto, se_abrio


#-EJECUTAR AGREGAR-

#PRE: el archivo debe existir
#POST: busca el id mas grande que haya en el archivo, le suma uno y lo devuelve
#      Devolviendo el id mas grande posible sin usar, considerando que el id más chico posible es 1
def id_mas_grande_sin_usar(archivo):
    reservas = open(archivo)  
    id_mayor = 0  
    lector = csv.reader(reservas, delimiter = DELIMITER)
    for reserva in lector:
        if(int(reserva[v.ID]) > id_mayor):
            id_mayor=int(reserva[v.ID])
    reservas.close()
    return id_mayor+1

#PRE: 
#POST: Lleva a cabo el comando agregar, agregando una linea en el archivo dado
#      la cual sera completada con un id calculado y los datos propocionados
def ejecutar_funcion_agregar(mensaje, archivo):

    id = str(id_mas_grande_sin_usar(archivo))
    nombre = mensaje[v.POSICION_NOMBRE]
    cantidad_personas = mensaje[v.POSICION_CANTIDAD_PERSONAS]
    horario = mensaje[v.POSICION_HORARIO]
    ubicacion = mensaje [v.POSICION_UBICACION]
    reserva_agregar = [id , nombre , cantidad_personas , horario , ubicacion]

    se_agrego = False
    reservas, se_abrio = abrir_archivo( archivo , MODO_AGREGAR )
    if se_abrio:
        escritor = csv.writer(reservas, delimiter = DELIMITER)
        escritor.writerow(reserva_agregar)
        reservas.close()
        se_agrego = True
    if se_agrego:
        print("Se agrego correctamente la reserva")




#-EJECUTAR MODIFICACION-

#PRE: 
#POST: realiza el cambio pedido en la linea, segun el campo indicado
def cambiar_linea ( nueva_linea , modificacion):
    campo, cambio  = modificacion[v.CAMPO_PEDIDO], modificacion[v.CAMBIO_PEDIDO]
    if campo == v.CAMPOS[v.NOMBRE]:
        nueva_linea[v.NOMBRE] = cambio
        
    elif campo == v.CAMPOS[v.CANT_PERSONAS]:
        nueva_linea[v.CANT_PERSONAS] = cambio

    elif campo == v.CAMPOS[v.HORARIO]:
        nueva_linea[v.HORARIO] = cambio

    elif campo == v.CAMPOS[v.UBICACION]:
        nueva_linea[v.UBICACION] = cambio

    return nueva_linea

#PRE: 
#POST: hace la modificacion pedida en la reserva del ID dado
#      si se dio un ID que no existe, lo informa y no hace ningun cambio
def realizar_modificacion( id_modificar , modificacion, archivo) :
    reservas, abrio_reservas = abrir_archivo( archivo , MODO_LECTURA)
    aux, abrio_aux = abrir_archivo( AUXILIAR , MODO_ESCRITURA )
    if abrio_reservas and abrio_aux:
        lector = csv.reader(reservas , delimiter = DELIMITER)
        escritor = csv.writer(aux , delimiter = DELIMITER)
        modificado = False
        for reserva in lector:
            if reserva[v.ID] == id_modificar:
                linea_nueva=cambiar_linea(reserva , modificacion)
                escritor.writerow(linea_nueva)
                modificado = True
            else: 
                escritor.writerow(reserva)
        os.rename(AUXILIAR, archivo)
        reservas.close()
        aux.close()
        if modificado:
            print("Se realizo correctamente la modificacion")
        else:
            print("No se encontro el ID dado")
    else:
        reservas.close()
        aux.close()

#PRE: 
#POST: solicita la informacion necesaria para la modificacion hasta que sea ingresada correctamente
#      una vez dada, ejecuta la modificacion
def ejecutar_funcion_modificar( id_modificar , archivo ):
    validado = False
    while not validado:
        modificacion = input("Ingrese campo y modificacion\n").split()   
        validado=v.validar_modificacion(modificacion) 
    if validado:      
        realizar_modificacion( id_modificar , modificacion, archivo)


#-EJECUTAR ELIMINAR-

#PRE: 
#POST: Elimina de la lista la linea(reserva) que tenga el ID indicado
#      Si no encuentra el id lo comunica
def ejecutar_funcion_eliminar(id_eliminar , archivo):
    reservas, abrio_reservas = abrir_archivo( archivo , MODO_LECTURA)
    aux, abrio_aux = abrir_archivo( AUXILIAR , MODO_ESCRITURA )
    if abrio_reservas and abrio_aux:
        lector = csv.reader(reservas , delimiter = DELIMITER)
        escritor = csv.writer(aux , delimiter = DELIMITER)
        se_borro = False
        for reserva in lector:
            if reserva[v.ID]!= id_eliminar:
                escritor.writerow(reserva)
            else:
                se_borro = True
        os.rename(AUXILIAR, archivo)
        reservas.close()
        aux.close()
        if se_borro:
            print(f"Se elimino correctamente la reserva {id_eliminar}")
        else:
            print("No se encontro el ID")
    else:
        reservas.close()
        aux.close()


#-EJECUTAR LISTAR-

#PRE: La lista debe tener el formato del .csv (id;nombre;cant_perosnas;HH:MM;ubicacion)
#POST: Muestra los datos de la lista(reserva)
def mostrar_reserva( reserva ):
    print(f"ID {reserva[v.ID]}\n"
          f"Nombre: {reserva[v.NOMBRE]}\n"
          f"Cantidad de personas: {reserva[v.CANT_PERSONAS]}\n"
          f"Hora: {reserva[v.HORARIO]}\n"
          f"Ubicacion: {reserva[v.UBICACION]}\n" )
    
#PRE:
#POST: Imprime en pantalla las reservas que cumplan la condicion.
#      Si no se especifico un rango, se imprimen todas 
#      Si se especifico, la condicion cambia y solo se imprimen las reservas dentro del rango
def ejecutar_funcion_listar( mensaje , largo, archivo):   
    reservas, se_abrio = abrir_archivo( archivo , MODO_LECTURA)
    if se_abrio:
        lector = csv.reader(reservas, delimiter = DELIMITER)
        condicion = (largo == v.CANTIDAD_DE_ARGUMENTOS_LISTAR_ENTERO)
        for reserva in lector:            
            if largo == v.CANTIDAD_DE_ARGUMENTOS_LISTAR_PARAMETRIZADO:
                minimo, maximo = int(mensaje[v.RANGO_MIN_INDICADO]) , int(mensaje[v.RANGO_MAX_INDICADO])
                condicion = int(reserva[v.ID]) >= minimo and int(reserva[v.ID]) <= maximo
            if condicion: 
                mostrar_reserva( reserva )
        reservas.close()


#-PRINCIPAL-
#PRE:
#POST: Reconoce y ejecuta el comando 
def realizar_comando(mensaje, comando, archivo):
    if comando == v.FUNCION_AGREGRAR:
        ejecutar_funcion_agregar(mensaje , archivo)

    elif comando == v.FUNCION_MODIFICAR:
        ejecutar_funcion_modificar(mensaje[ID_INDICADO], archivo)

    elif comando == v.FUNCION_ELIMINAR:
        ejecutar_funcion_eliminar(mensaje[ID_INDICADO], archivo)
        pass

    elif comando == v.FUNCION_LISTAR:
        ejecutar_funcion_listar(mensaje , len(mensaje), archivo)