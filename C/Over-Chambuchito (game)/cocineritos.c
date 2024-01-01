#include <stdio.h>
#include "cocineritos.h"
#include <stdlib.h> // Para usar rand
#include <time.h> // Para obtener una semilla desde el reloj

//##CONSTANTES##//

/* Constantes generales(se usan para herramientas y a lo largo de todo el programa) */
#define MAXIMO_MATRIZ 21
#define MINIMO_MATRIZ 0
#define MITAD_MATRIZ 10
#define CUADRANTE_STITCH 1
#define CUADRANTE_REUBEN 2
#define MINIMA_FILA_CUADRANTE_STITCH 1
#define MINIMA_FILA_CUADRANTE_REUBEN 11
#define RANGO_FILA_CUADRANTE_STITCH 8
#define RANGO_FILA_CUADRANTE_REUBEN 8
#define MINIMA_COLUMNA_CUADRANTE_STITCH 1
#define MINIMA_COLUMNA_CUADRANTE_REUBEN 1
#define RANGO_COLUMNA_CUADRANTE_STITCH 18
#define RANGO_COLUMNA_CUADRANTE_REUBEN 18
const int MAX_CANTIDAD_CUCHILLOS = 2 ;
const int MAX_CANTIDAD_HORNOS = 2 ;
const int MAX_CANTIDAD_AGUJEROS_TOTAL = 20 ;
const int MAX_CANTIDAD_AGUJEROS_POR_CUADRANTE = 10 ;
#define CANTIDAD_INGREDIENTES_ENSALADA 2
#define CANTIDAD_INGREDIENTES_PIZZA 3
#define CANTIDAD_INGREDIENTES_HAMBURGUESA 4
#define CANTIDAD_INGREDIENTES_SANDWICH 6
const int PRECIO_ENSALADA_PIZZA_HAMBURGUESA = 100 ;
const int PRECIO_ENSALADA_PIZZA_HAMBURGUESA_SANDWICH = 150 ;




//Constantes para jugadas//
#define LETRA_MOVER_ARRIBA 'W'
#define LETRA_MOVER_ABAJO 'S'
#define LETRA_MOVER_IZQUIERDA 'A'
#define LETRA_MOVER_DERECHA 'D'
#define LETRA_CAMBIAR_PERSONAJE 'X'
#define LETRA_USAR_MATAFUEGOS 'M'
#define LETRA_USAR_CUCHILLO 'C'
#define LETRA_USAR_HORNO 'H'
#define LETRA_AGARRAR_SOLTAR_INGREDIENTE 'R'
#define LETRA_USAR_MESA 'T'
const int DISTANCIA_PARA_COCINAR = 1 ;
const int DISTANCIA_PARA_USAR_MESA = 1 ;
const int DISTANCIA_PARA_USAR_MATAFUEGOS = 2 ;
const int VALOR_MOVIMIENTO_POSITIVO = 1 ;
const int VALOR_MOVIMIENTO_NEGATIVO = (-1) ;
const int VALOR_MOVIMIENTO_NULO = 0 ;




//constantes para personajes, objetos y comidas//
#define EMOJI_VACIO "\U0001f538"
#define EMOJI_STITCH "\U0001f436"
#define EMOJI_REUBEN "\U0001f981"
#define EMOJI_PARED "\U0001f9f1"
#define EMOJI_FUEGO "\U0001f525"
#define EMOJI_AGUJERO "\U0001f573\uFE0F"
#define EMOJI_MATAFUEGO "\U0001f9ef"
#define EMOJI_CUCHILLO "\U0001f52a"
#define EMOJI_HORNO "\U0001f39b\uFE0F"
#define EMOJI_SALIDA "\U0001f6aa"
#define EMOJI_MESA "\U0001f532"
#define EMOJI_LECHUGA " \U0001f96c"
#define EMOJI_TOMATE "\U0001f345"
#define EMOJI_MILANESA "\U0001f357"
#define EMOJI_CARNE "\U0001f356"
#define EMOJI_PAN "\U0001f35e"
#define EMOJI_JAMON "\U0001f953"
#define EMOJI_QUESO "\U0001f9c0"
#define EMOJI_MASA "\U0001f963"
#define LETRA_STITCH 'S'
#define LETRA_REUBEN 'R'
const char LUGAR_VACIO = ' ' ;
const char PARED = '#' ;
const char SALIDA = 'P' ;
const char MESA = '_' ;
const char AGUJEROS = 'A' ;
const char FUEGO =  'F' ;
const char HORNOS = 'H' ;
const char CUCHILLOS = 'C' ;
const char MATAFUEGOS = 'M' ;
#define LECHUGA 'L'
#define TOMATE 'T'
#define MASA 'O'
#define QUESO 'Q'
#define JAMON 'J'
#define CARNE 'B'
#define PAN 'N'
#define MILANESA 'I'
const char ENSALADA='E';
const char PIZZA = 'P';
const char HAMBURGUESA='H';
const char SANDWICH='S';
const char NINGUN_OBJETO='V';




//constantes para estado juego//
const static int JUEGO_GANADO = 1 ;
const static int SIGUE_JUEGO = 0 ;
const static int JUEGO_PERDIDO = -1 ;







/* FUNCIONES "HERRAMIENTAS" */




// PRE-CONDICIONES: -
//POST-CONDICIONES: Devuelve un numero aleatorio, entre el minimo y rango ingresados.
int numero_random( int minimo , int rango ) {
    return rand() % ( rango ) + minimo ;
}



// PRE-CONDICIONES:-
//POST-CONDICIONES: Compara dos coordenadas dadas, true si son iguales, false si no lo son
bool comparar_coordenas( coordenada_t coordenada_comparada_1 , coordenada_t coordenada_comparada_2 ) {
    return coordenada_comparada_1.fil == coordenada_comparada_2.fil &&
           coordenada_comparada_1.col == coordenada_comparada_2.col ;        
}



/* PRE-CONDICIONES: -Se le dara el tope_comida sin modificar cuando NO haya alguna 
                    comida en proceso (cuando se inicializa o carga una nueva comida)

                    -Se le dara el tope_comida-1 cuando haya la comida ya este incializada (generar un personaje, en mitad de un nivel para el fuego, etc)*/

//POST-CONDICIONES: Compara si la coordenada ingresada ya esta asignada para algun objeto/ingrediente/personaje/mesa/salida.
bool esta_usada_coordenada( coordenada_t coordenada_nueva , juego_t juego , int tope_comida ) {
    bool fue_usada = false ;
    for( int i = 0 ; i < juego.tope_obstaculos ; i++ ){
        if ( comparar_coordenas ( coordenada_nueva , juego.obstaculos[i].posicion ) ) fue_usada = true;
    }
    for( int i = 0; i < juego.tope_herramientas ; i++ ){
        if ( comparar_coordenas ( coordenada_nueva , juego.herramientas[i].posicion ) ) fue_usada = true;
    }
    for( int j = 0 ; j < juego.comida[tope_comida].tope_ingredientes ; j++ ){
        if ( comparar_coordenas ( coordenada_nueva , juego.comida[tope_comida].ingrediente[j].posicion ) ) fue_usada = true;        
    }
    if ( comparar_coordenas (coordenada_nueva , juego.stitch.posicion ) || comparar_coordenas ( coordenada_nueva , juego.reuben.posicion ) ) fue_usada = true;
    if ( comparar_coordenas ( coordenada_nueva , juego.salida ) ) fue_usada = true ;

    return fue_usada ;
}



// PRE-CONDICIONES: Se le dara un entero que indicara en que cuadrante debe generar la coordenada. CUADRANTE_STITCH = 1
//                                                                                                 CUADRANTE_REUBEN = 2
/*                  -Se le dara el tope_comida sin modificar cuando NO haya alguna 
                    comida en proceso (cuando se inicializa o carga una nueva comida)

                    -Se le dara el tope_comida-1 cuando haya una comida en proceso (mitad de un nivel)*/

//POST-CONDICIONES: Generara una nueva coordenada aleatoria sin usar en el cuadrante indicado. 
coordenada_t generar_coordenada_aleatoria( int cuadrante, juego_t juego, int tope_comida ) {
    coordenada_t coordenada_aleatoria;
    do{
    if ( cuadrante == CUADRANTE_STITCH ) {
    coordenada_aleatoria.fil = numero_random ( MINIMA_FILA_CUADRANTE_STITCH , RANGO_FILA_CUADRANTE_STITCH ) ;
    coordenada_aleatoria.col = numero_random ( MINIMA_COLUMNA_CUADRANTE_STITCH , RANGO_COLUMNA_CUADRANTE_STITCH ) ;

    } else if( cuadrante == CUADRANTE_REUBEN ) {
    coordenada_aleatoria.fil = numero_random ( MINIMA_FILA_CUADRANTE_REUBEN , RANGO_FILA_CUADRANTE_REUBEN ) ; 
    coordenada_aleatoria.col = numero_random ( MINIMA_COLUMNA_CUADRANTE_REUBEN , RANGO_COLUMNA_CUADRANTE_REUBEN ) ;
        }

    }while( esta_usada_coordenada ( coordenada_aleatoria , juego, tope_comida ) ) ;
    return coordenada_aleatoria ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Busca y devuelve el tipo(letra) del obstaculo/herramienta/ingrediente que este en la misma posicion que el personaje indicado.
char objeto_debajo ( juego_t juego , personaje_t personaje ) {
    char objeto_debajo = LUGAR_VACIO ;
    for( int i = 0 ; i < juego.tope_obstaculos ; i++ ){
        if ( comparar_coordenas ( personaje.posicion , juego.obstaculos[i].posicion ) ) {
            objeto_debajo = juego.obstaculos[i].tipo ;
        }
    }   
    for( int i = 0 ; i < juego.tope_herramientas ; i++ ){
        if ( comparar_coordenas ( personaje.posicion , juego.herramientas[i].posicion ) ) {
            objeto_debajo = juego.herramientas[i].tipo ;
        }
    }
    for( int j = 0 ; j < juego.comida[juego.tope_comida-1].tope_ingredientes ; j++ ) {
        if ( comparar_coordenas ( personaje.posicion , juego.comida[juego.tope_comida-1].ingrediente[j].posicion ) ) {
            objeto_debajo = juego.comida[juego.tope_comida-1].ingrediente[j].tipo ;        
        }   
    }
    if (comparar_coordenas( personaje.posicion , juego.salida)){
        objeto_debajo = SALIDA;
    }

    return objeto_debajo ;

}



// PRE-CONDICIONES: Que el personaje tenga un ingrediente en mano.
//POST-CONDICIONES: Busca y devuelve el entero que indica la posicion de vector del ingrediente que tiene en la mano el personaje indicado.
//                    En caso de no encontrarlo devuelve -1 (posisiocn de vector invalida)
int buscar_numero_ingrediente_en_mano( juego_t juego , personaje_t personaje ) {
    int numero_ingrediente = -1 ;
    for( int i = 0 ; i < juego.comida[juego.tope_comida-1].tope_ingredientes ; i++ ) {
            if( juego.comida[juego.tope_comida-1].ingrediente[i].tipo == personaje.objeto_en_mano ) numero_ingrediente = i ;

    }
    return numero_ingrediente ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Devuelve el tipo(letra) del objeto en mano del personaje activo 
char objeto_en_mano_personaje_activo (juego_t juego){
    char objeto_actual = ' ';
    if(juego.personaje_activo == LETRA_STITCH){
        objeto_actual = juego.stitch.objeto_en_mano;

    } else if(juego.personaje_activo == LETRA_REUBEN){
        objeto_actual = juego.reuben.objeto_en_mano;
    }

    return  objeto_actual;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Calcula la distancia manhattan entre dos coordenadas dadas. 
int calcular_distancia_manhattan ( coordenada_t coordenada_1 , coordenada_t coordenada_2 ) {
    return abs( ( coordenada_2.fil - coordenada_1.fil )  +  ( coordenada_2.col - coordenada_1.col ) ) ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Recorre los obstaculos para verificar si hay fuego en el juego.
//                  Devuelve si hay fuego o no.
bool hay_fuego ( juego_t juego ) {
    bool se_encontro = false ;
    for( int i = 0 ; i<juego.tope_obstaculos ; i++){
        if ( juego.obstaculos[ i ].tipo == FUEGO){
            se_encontro = true ;
        }
    }
    return se_encontro ;
}



// PRE-CONDICIONES: Debe haber fuego en el juego
//POST-CONDICIONES: Devuelve el entero que indica la fila del fuego. Si no lo encuentra, devuelve una fila invalida
int fila_del_fuego( juego_t juego ){
    int fila_fuego = -1 ;
    for( int i = 0 ; i<juego.tope_obstaculos ; i++){
        if ( juego.obstaculos[ i ].tipo == FUEGO){
            fila_fuego = juego.obstaculos[ i ].posicion.fil;
        }
    }
    return fila_fuego;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Devuelve si hay o no hay fuego en el cuadrante del personaje activo.
bool hay_fuego_cuadrante_personaje_activo ( juego_t juego , char personaje_activo ) {
    bool cuadrante = false ;
    if ( hay_fuego ( juego ) && fila_del_fuego(juego) < MITAD_MATRIZ && personaje_activo == LETRA_STITCH ) {
        cuadrante = true;
    }
    if ( hay_fuego ( juego ) && fila_del_fuego(juego) > MITAD_MATRIZ && personaje_activo == LETRA_REUBEN ) {
        cuadrante = true ;
    }
return cuadrante;

}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Devuelve el tipo(letra) del ingrediente que hay en la mesa.
//                  Si no hay ningun ingrediente, devuelve que no hay ningun objeto.
char ingrediente_en_mesa ( juego_t juego ) {
    char tipo_ingrediente_en_mesa = NINGUN_OBJETO ;
    for( int i = 0 ; i < juego.comida[(juego.tope_comida)-1].tope_ingredientes ; i++ ) {
        if(comparar_coordenas ( juego.mesa , juego.comida[(juego.tope_comida)-1].ingrediente[i].posicion ) ) {
            tipo_ingrediente_en_mesa = juego.comida[(juego.tope_comida)-1].ingrediente[i].tipo ;
        }
    }
    return tipo_ingrediente_en_mesa ;
}






/* PROCEDIMIENTOS Y FUNCIONES PARA INCIALIZAR */




// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicializa todos los topes.
void inicializar_topes ( juego_t* juego ) {
    juego->tope_paredes = 0 ;
    juego->tope_herramientas = 0 ;
    juego->tope_comida = 0 ;
    juego->tope_obstaculos = 0 ;
    juego->tope_comida_lista = 0;
    
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Verifica si la fila y columna dada cumplen con la condicion de pared
//                  (la condiciones tener coordenada de pared, porque las coordenadas de pared son fijas)
bool condicion_pared( int fila , int columna ) {
    return ( fila == MINIMO_MATRIZ  ||  columna == MINIMO_MATRIZ  ||  fila == (MAXIMO_MATRIZ-1)  ||  columna == (MAXIMO_MATRIZ-1)  ||  ( fila == MITAD_MATRIZ && columna != MITAD_MATRIZ ) );
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicializa las paredes, dandole una coordenada a cada una.
void inicializar_paredes ( coordenada_t paredes [ MAX_PAREDES ] , int *tope_paredes ) {
    *tope_paredes = 0 ;
    for ( int i = MINIMO_MATRIZ ; i < MAXIMO_MATRIZ ; i++ ) {
        for( int j = MINIMO_MATRIZ ; j < MAXIMO_MATRIZ ; j++ ) {
            if ( condicion_pared ( i , j ) ) {
                paredes[*tope_paredes].fil = i ;
                paredes[*tope_paredes].col = j ;
                (*tope_paredes)++ ;
            }
        }
    }
}



// PRE-CONDICIONES: Se le debe indicar el cuadrante
//POST-CONDICIONES: Inicializa los agujeros en el cuadrante indicado, dandoles coordenada y su tipo correspondiente.
void inicializar_agujeros_cuadrante_elegido ( juego_t* juego , int cuadrante_elegido ) {
    for ( int i = 0 ; i < MAX_CANTIDAD_AGUJEROS_POR_CUADRANTE ; i++ ) {
        juego->obstaculos[juego->tope_obstaculos].posicion = generar_coordenada_aleatoria ( cuadrante_elegido , (*juego) , juego->tope_comida ) ;
        juego->obstaculos[juego->tope_obstaculos].tipo = AGUJEROS ;
        juego->tope_obstaculos++ ;
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicializa los agujeros del mapa
void inicializar_agujeros ( juego_t *juego ) {
    inicializar_agujeros_cuadrante_elegido ( & (*juego) , CUADRANTE_STITCH ) ;
    inicializar_agujeros_cuadrante_elegido ( &(*juego) , CUADRANTE_REUBEN ) ;

}



// PRE-CONDICIONES: Se le debe indicar el tipo, cantidad y el cuadrante donde estara la herramienta
//POST-CONDICIONES: Inicializa herramientas con los datos dados
void inicializar_herramientas_elegida ( juego_t *juego , int cuadrante_elegido , char tipo_herramienta_elegida , int cantidad_herramientas_elegida ) {
    for( int i = 0 ; i < cantidad_herramientas_elegida ; i++ ) {
        juego->herramientas[juego->tope_herramientas].posicion = generar_coordenada_aleatoria ( cuadrante_elegido , (*juego) , juego->tope_comida ) ;
        juego->herramientas[juego->tope_herramientas].tipo = tipo_herramienta_elegida ;
        juego->tope_herramientas++ ;
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicaliza los cuchillos y hornos.
void inicializar_herramientas( juego_t *juego ) {
    inicializar_herramientas_elegida( &(*juego) , CUADRANTE_STITCH , CUCHILLOS , MAX_CANTIDAD_CUCHILLOS ) ;
    inicializar_herramientas_elegida( &(*juego) , CUADRANTE_REUBEN , HORNOS , MAX_CANTIDAD_HORNOS ) ;
}



// PRE-CONDICIONES: Se le debe indicar el tipo
//POST-CONDICIONES: Inicializa el ingrediente con los datos dados.
void inicializar_ingrediente_elegido ( juego_t *juego , int cuadrante_elegido , char tipo_ingrediente_elegido ) {
    juego->comida[ (*juego).tope_comida ].ingrediente[ juego->comida[ (*juego).tope_comida ].tope_ingredientes ]. posicion = generar_coordenada_aleatoria ( cuadrante_elegido , (*juego) , juego->tope_comida);
    juego->comida[(*juego).tope_comida].ingrediente[juego->comida[(*juego).tope_comida].tope_ingredientes ].tipo = tipo_ingrediente_elegido ;
    juego->comida[(*juego).tope_comida].ingrediente[juego->comida[(*juego).tope_comida].tope_ingredientes ].esta_cortado = false ;
    juego->comida[(*juego).tope_comida].ingrediente[juego->comida[(*juego).tope_comida].tope_ingredientes ].esta_cocinado = false ;
    (juego->comida[(*juego).tope_comida].tope_ingredientes)++ ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicializa la Ensalada, dandole tipo, coordenada e inicializando sus ingrediente. Y la asigna como comida actual 
void inicializar_ensalada ( juego_t *juego ) {
    juego->comida_actual = ENSALADA;
    juego->comida[(*juego).tope_comida].tipo = ENSALADA ;
    (juego->comida[(*juego).tope_comida].tope_ingredientes) = 0 ;
    inicializar_ingrediente_elegido ( &(*juego) , CUADRANTE_STITCH , LECHUGA ) ;
    inicializar_ingrediente_elegido ( &(*juego) , CUADRANTE_STITCH , TOMATE ) ;
    (juego->tope_comida)++ ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicializa la Pizza, dandole tipo, coordenada e inicializando sus ingrediente. Y la asigna como comida actual 
void inicializar_pizza ( juego_t *juego ) {
    juego->comida_actual = PIZZA ;
    juego->comida[ (*juego).tope_comida ].tipo = PIZZA ;
    (juego->comida[ (*juego).tope_comida ].tope_ingredientes) = 0 ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_REUBEN , MASA ) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , QUESO) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , JAMON) ;
    (juego->tope_comida)++ ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicializa la Hamburguesa, dandole tipo, coordenada e inicializando sus ingrediente. Y la asigna como comida actual 
void inicializar_hamburguesa(juego_t *juego){
    juego->comida_actual = HAMBURGUESA ;
    juego->comida[(*juego).tope_comida].tipo = HAMBURGUESA ;
    (juego->comida[(*juego).tope_comida].tope_ingredientes) = 0 ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , PAN ) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , LECHUGA ) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , TOMATE ) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_REUBEN , CARNE ) ;
    (juego->tope_comida)++ ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicializa el sandwich, dandole tipo, coordenada e inicializando sus ingrediente. Y lo asigna como comida actual 
void inicializar_sandwich(juego_t *juego){
    juego->comida_actual = SANDWICH ;
    juego->comida[(*juego).tope_comida].tipo = SANDWICH ;
    (juego->comida[(*juego).tope_comida].tope_ingredientes) = 0 ;
 
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , PAN ) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , LECHUGA ) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , TOMATE ) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_REUBEN , MILANESA ) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , QUESO) ;
    inicializar_ingrediente_elegido( &(*juego) , CUADRANTE_STITCH , JAMON) ;
    (juego->tope_comida)++ ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicializa el personaje indicado, dandole una coordenada correspondiente, tipo y objeto_en_mano.
void inicializar_personaje_elegido( juego_t *juego , personaje_t *personaje , char personaje_elegido , int cuadrante_elegido ) {
    (*personaje).posicion = generar_coordenada_aleatoria(cuadrante_elegido, (*juego), juego->tope_comida-1) ;
    (*personaje).tipo = personaje_elegido ;
    (*personaje).objeto_en_mano = NINGUN_OBJETO ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Inicializa los dos personajes del juego, ademas de indicar cual empieza como personaje_activo
void inicializar_personajes( juego_t *juego ) {
    juego->personaje_activo = LETRA_STITCH ;
    inicializar_personaje_elegido (&(*juego) , &(*juego).stitch , LETRA_STITCH , CUADRANTE_STITCH ) ;
    inicializar_personaje_elegido (&(*juego) , &(*juego).reuben , LETRA_REUBEN , CUADRANTE_REUBEN ) ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Le asigna coordenada a la mesa.
void inicializar_mesa ( juego_t *juego ) {
    juego->mesa.fil = MITAD_MATRIZ ;
    juego->mesa.col = MITAD_MATRIZ ;
}






/* PROCEDIMIENTOS Y FUNCIONES PARA IMPRIMIR MATRIZ */




// PRE-CONDICIONES: -
//POST-CONDICIONES: Carga la matriz de lugares vacios.
void cargar_matriz ( char matriz[ MAXIMO_MATRIZ ][ MAXIMO_MATRIZ ] ) {
    for( int i = 0 ; i < MAXIMO_MATRIZ ; i++ ) {
        for( int j = 0 ; j < MAXIMO_MATRIZ ; j++ ) {
           matriz[ i ][ j ] = LUGAR_VACIO ; 
        }
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Carga las paredes en la matriz, con su convencion y coordenadas correspondientes
void cargar_paredes ( juego_t juego , char matriz[ MAXIMO_MATRIZ ][ MAXIMO_MATRIZ ] ) {
    for( int i = 0 ; i < juego.tope_paredes ; i++ ) {
        matriz[ juego.paredes[ i ].fil ][ juego.paredes[ i ].col ] = PARED;
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Carga los obstaculos en la matriz, con su convencion y coordenadas correspondientes
void cargar_obstaculos ( juego_t juego , char matriz[ MAXIMO_MATRIZ ][ MAXIMO_MATRIZ ] ) {

    for( int i = 0 ; i < juego.tope_obstaculos ; i++ ) {
        if(juego.obstaculos[ i ].tipo == AGUJEROS){
            matriz[ juego.obstaculos[ i ].posicion.fil ][ juego.obstaculos[ i ].posicion.col ] = AGUJEROS ;
        }
        if ( hay_fuego (juego) && juego.obstaculos[ i ].tipo == FUEGO ) {
            matriz[ juego.obstaculos[ i ].posicion.fil ][ juego.obstaculos[ i ].posicion.col ] = FUEGO ;
        }
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Carga las herramientas en la matriz, con su convencion y coordenadas correspondientes
void cargar_herramientas ( juego_t juego , char matriz[ MAXIMO_MATRIZ ][ MAXIMO_MATRIZ ] ) {

    for ( int i = 0 ; i < juego.tope_herramientas ; i++ ) {

        if ( juego.herramientas[ i ].tipo == CUCHILLOS ) {
            matriz[ juego.herramientas[ i ].posicion.fil ][ juego.herramientas[ i ].posicion.col ] = CUCHILLOS ; 

        }else if ( juego.herramientas[ i ].tipo == HORNOS ) {
            matriz[ juego.herramientas[ i ].posicion.fil ][ juego.herramientas[ i ].posicion.col ] = HORNOS ;
            
        }else if (juego.herramientas[ i ].tipo == MATAFUEGOS 
                  && juego.stitch.objeto_en_mano != MATAFUEGOS && juego.reuben.objeto_en_mano != MATAFUEGOS  ) {
                        matriz[ juego.herramientas[ i ].posicion.fil ][ juego.herramientas[ i ].posicion.col ] = MATAFUEGOS ;
                }
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Carga los ingredientes en la matriz, con su convencion y coordenadas correspondientes. 
//                  Si el personaje tiene un ingrediente en mano, el/los ingredientes que sean de mismo tipo no se cargan en la matriz.
//                  (se hacen "invisibles")
void cargar_ingredientes ( juego_t juego , char matriz[ MAXIMO_MATRIZ ][ MAXIMO_MATRIZ ] ) {

    for ( int j = 0 ; j < juego.comida[ juego.tope_comida-1 ].tope_ingredientes ; j++ ) {

        if ( juego.stitch.objeto_en_mano != juego.comida[ juego.tope_comida-1 ].ingrediente[ j ].tipo   
            &&   juego.reuben.objeto_en_mano != juego.comida[ juego.tope_comida-1 ].ingrediente[ j ].tipo ) {

            matriz[ juego.comida[ juego.tope_comida-1 ].ingrediente[ j ].posicion.fil ][ juego.comida[ juego.tope_comida-1 ].ingrediente[ j ].posicion.col ] = 
            juego.comida[ juego.tope_comida-1 ].ingrediente[ j ].tipo ;
        }
    }    
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Carga los personajes en la matriz, con su convencion y coordenadas correspondientes
void cargar_personajes ( juego_t juego , char matriz[ MAXIMO_MATRIZ ][ MAXIMO_MATRIZ ] ) {
    matriz[ juego.stitch.posicion.fil ][ juego.stitch.posicion.col ] = LETRA_STITCH ;
    matriz[ juego.reuben.posicion.fil ][ juego.reuben.posicion.col ] = LETRA_REUBEN ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Carga la mesa y la salida en la matriz, con su convencion y coordenadas correspondientes
void cargar_mesa_y_salida ( juego_t juego , char matriz[ MAXIMO_MATRIZ ][ MAXIMO_MATRIZ ] ) {
    matriz[ juego.salida.fil ] [juego.salida.col ] = SALIDA ;
    matriz[ juego.mesa.fil ][ juego.mesa.col ] = MESA ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Muestra la matriz en pantallla
/*void mostrar_matriz ( char matriz[ MAXIMO_MATRIZ ][ MAXIMO_MATRIZ ] ) {

    for( int i = 0 ; i < MAXIMO_MATRIZ ; i++ ) {
            for( int j = 0 ; j < MAXIMO_MATRIZ ; j++ ) {
            printf ( " %c" , matriz[ i ][ j ] ) ; 
            }
            printf ( "\n" ) ;
        }
}*/
void mostrar_matriz(char matriz[MAXIMO_MATRIZ][MAXIMO_MATRIZ]){
    for (int i = 0; i < MAXIMO_MATRIZ; i++){
        printf("          ");
        for (int j = 0; j < MAXIMO_MATRIZ; j++){
            if (matriz[i][j] == PARED) printf(" %s", EMOJI_PARED);
            else if (matriz[i][j] == LUGAR_VACIO) printf(" %s", EMOJI_VACIO);
            else if (matriz[i][j] == MESA) printf(" %s", EMOJI_MESA);
            else if (matriz[i][j] == AGUJEROS) printf(" %s ", EMOJI_AGUJERO);
            else if (matriz[i][j] == HORNOS) printf(" %s ", EMOJI_HORNO);
            else if (matriz[i][j] == CUCHILLOS) printf(" %s", EMOJI_CUCHILLO);
            else if (matriz[i][j] == LETRA_STITCH) printf(" %s", EMOJI_STITCH);
            else if (matriz[i][j] == LETRA_REUBEN) printf(" %s", EMOJI_REUBEN);
            else if (matriz[i][j] == TOMATE) printf(" %s", EMOJI_TOMATE);
            else if (matriz[i][j] == LECHUGA) printf("%s", EMOJI_LECHUGA);
            else if (matriz[i][j] == PAN) printf(" %s", EMOJI_PAN);
            else if (matriz[i][j] == CARNE) printf(" %s", EMOJI_CARNE);
            else if (matriz[i][j] == FUEGO) printf(" %s", EMOJI_FUEGO);
            else if (matriz[i][j] == MATAFUEGOS) printf(" %s", EMOJI_MATAFUEGO);
            else if (matriz[i][j] == SALIDA) printf(" %s", EMOJI_SALIDA);
            else if (matriz[i][j] == JAMON) printf(" %s", EMOJI_JAMON);
            else if (matriz[i][j] == QUESO) printf(" %s", EMOJI_QUESO);
            else if (matriz[i][j] == MASA) printf(" %s", EMOJI_MASA);
            else if (matriz[i][j] == MILANESA) printf(" %s", EMOJI_MILANESA);
        }
        printf("\n");
    }
}

void mostrar_info_personaje(juego_t juego, personaje_t personaje){
    if (personaje.tipo == LETRA_STITCH) printf("Personaje activo: %s \n", EMOJI_STITCH);
    else if (personaje.tipo == LETRA_REUBEN) printf("Personaje activo: %s \n", EMOJI_REUBEN);
    if (personaje.objeto_en_mano == TOMATE) printf("Objeto en mano: %s \n", EMOJI_TOMATE);
    else if (personaje.objeto_en_mano == LECHUGA) printf("Objeto en mano: %s \n", EMOJI_LECHUGA);
    else if (personaje.objeto_en_mano == PAN) printf("Objeto en mano: %s \n", EMOJI_PAN);
    else if (personaje.objeto_en_mano == JAMON) printf("Objeto en mano: %s \n", EMOJI_JAMON);
    else if (personaje.objeto_en_mano == QUESO) printf("Objeto en mano: %s \n", EMOJI_QUESO);
    else if (personaje.objeto_en_mano == MASA) printf("Objeto en mano: %s \n", EMOJI_MASA);
    else if (personaje.objeto_en_mano == MILANESA) printf("Objeto en mano: %s \n", EMOJI_MILANESA);
    else if (personaje.objeto_en_mano == MATAFUEGOS) printf("Objeto en mano: %s \n", EMOJI_MATAFUEGO);
    for (int i = 0; i < juego.comida[juego.tope_comida-1].tope_ingredientes; i++){
        if(juego.comida[juego.tope_comida-1].ingrediente[i].tipo == personaje.objeto_en_mano){
            if (juego.comida[juego.tope_comida-1].ingrediente[i].esta_cortado){
                printf("          Esta cortado \n");
            } else {
                printf("          No esta cortado \n");
            }
            if (juego.comida[juego.tope_comida-1].ingrediente[i].esta_cocinado){
                printf("          Esta cocinado \n");
            } else {
                printf("          No esta cocinado \n");
            }
        }
    }
}



/* PROCEDIMIENTOS Y FUNCIONES PARA REALIZAR JUGADA */




// PRE-CONDICIONES: Los valores movimientos deben tener modulo 1
//POST-CONDICIONES: Verifica si el siguiente movimiento es posible, si el siguiente movimiento implica moverse hacia algun HORNO, PARED, FUEGO. No sera posible. 
bool verificar_siguiente_movimiento ( juego_t juego , personaje_t personaje , int valor_movimiento_fil , int valor_movimiento_col ) {
    bool es_posible = true ;
    coordenada_t siguiente_posicion_posible ;
    siguiente_posicion_posible.fil = personaje.posicion.fil + valor_movimiento_fil ;
    siguiente_posicion_posible.col = personaje.posicion.col + valor_movimiento_col ;

    for( int i = 0 ; i < juego.tope_paredes ; i++ ) {
        if( comparar_coordenas ( siguiente_posicion_posible , juego.paredes[ i ] )  ||  comparar_coordenas ( siguiente_posicion_posible , juego.mesa ) ) {
            es_posible = false ;
        }
    }

    for( int i = 0 ; i < juego.tope_herramientas ; i++ ) {
        if( comparar_coordenas ( siguiente_posicion_posible , juego.herramientas[ i ].posicion )  &&  juego.herramientas[ i ].tipo == HORNOS ) {
        es_posible = false ;
        }
    }   

    if ( hay_fuego ( juego ) && comparar_coordenas(siguiente_posicion_posible , juego.obstaculos[ juego.tope_obstaculos - 1].posicion ) ){
        es_posible = false ;
    }

    return es_posible ;
}



// PRE-CONDICIONES: Los valores movimientos deben tener modulo 1
//POST-CONDICIONES: Si es posible, realiza el movimiento con el personaje activo modificando su coordenada según los valores dados. 
//                  Suma 1 al contador movimientos si no hay fuego.
void mover_personaje ( juego_t *juego , personaje_t *personaje , int valor_movimiento_fil , int valor_movimiento_col ) {
    if(verificar_siguiente_movimiento ( (*juego) , (*personaje) , valor_movimiento_fil , valor_movimiento_col ) ) {
        ( personaje->posicion.fil ) += ( valor_movimiento_fil ) ;
        ( personaje->posicion.col ) += ( valor_movimiento_col ) ;
        if( !hay_fuego (*juego ) ){
            ( juego->movimientos )++ ;
        }
    }
    
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Cambia el personaje_activo
void cambiar_personaje ( juego_t *juego ) {
    if ( juego->personaje_activo == LETRA_STITCH ) {
        juego->personaje_activo = LETRA_REUBEN ;
    } else {
        juego->personaje_activo = LETRA_STITCH ;
    }
}



// PRE-CONDICIONES: No debe haber fuego en el cuadrante del personaje
//                  el personaje activo debe tener la mano vacia y debe compartir coordenada con un ingrediente (estar arriba)
//POST-CONDICIONES: El objeto en mano del personaje cambia al ingrediente que tiene misma coordenada que el
//                  (agarra el ingrediente que tiene debajo) El ingrediente no desaparece, solamente no se ve maś
void agarrar_ingrediente ( juego_t *juego , personaje_t *personaje ) {
    for ( int i = 0 ;  i < juego->comida[ juego->tope_comida - 1 ].tope_ingredientes ; i++ ) {
            if ( comparar_coordenas ( personaje->posicion , juego->comida[ juego->tope_comida - 1 ].ingrediente[ i ].posicion ) ) {
            personaje->objeto_en_mano = juego->comida[ juego->tope_comida - 1 ].ingrediente[ i ].tipo ;
        }
    }  
}


// PRE-CONDICIONES: El personaje activo debe tener un ingrediente en la mano y no debe compartir coordenada con nada, o solo con el ingrediente que tiene en mano
//                  (no estar arriba de nada o arriba del ingrediente que tiene)
//POST-CONDICIONES: El ingrediente que tiene en mano cambia de posicion, se le asigna la misma que el personaje 
//                  El objeto en mano del personaje cambia a ningun objeto
//                  (Suelta el ingrediente)
void soltar_ingrediente ( juego_t *juego , personaje_t *personaje ) {
    if( objeto_debajo( *juego , *personaje ) ==  LUGAR_VACIO  
        || objeto_debajo ( *juego , *personaje ) == personaje->objeto_en_mano ) {
            juego->comida[ juego->tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( *juego , *personaje ) ].posicion = personaje->posicion ;
            personaje->objeto_en_mano = NINGUN_OBJETO ;
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Se definira si se puede soltar o agarrar el ingrediente, y se ejecutara lo que corresponda
void agarrar_soltar_ingrediente ( juego_t *juego , personaje_t *personaje ) {
    if ( personaje->objeto_en_mano == NINGUN_OBJETO && !hay_fuego_cuadrante_personaje_activo ( *juego , personaje->tipo ) ) {
        agarrar_ingrediente ( &( *juego ) , &( *personaje ) ) ;
    }else if ( personaje->objeto_en_mano !=  MATAFUEGOS ) {
        soltar_ingrediente ( &( *juego ) , &( *personaje ) ) ;
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Verifica si el ingrediente es posible de cortar, según su tipo
bool se_puede_cortar ( char tipo_ingrediente ){
    bool se_puede = false ;
    switch ( tipo_ingrediente ) {
        case LECHUGA :
            se_puede = true ;            
            break ;

        case TOMATE :
            se_puede = true ;            
        break ;

        case QUESO :
            se_puede = true ;            
        break ;

        case JAMON :
            se_puede = true ;
        break ;

        case PAN :
            se_puede = true ;
        break ;

    }
    return se_puede ;
}


    
// PRE-CONDICIONES: El personaje activo debe ser stitch
//                  debe compartir coordenada con un cuchillo(estar arriba) y tener un ingrediente posible de cortar
//POST-CONDICIONES: El ingrediente que tenga el mismo tipo que el que tiene stitch en la mano, cambiara su booleano "esta cortado" a true 
//                  (Se corta el ingrediente que tiene stitch en la mano)
void usar_cuchillo ( juego_t *juego , personaje_t stitch ) {
    if( juego->personaje_activo == LETRA_STITCH
     && objeto_debajo ( *juego , stitch ) == CUCHILLOS 
     && stitch.objeto_en_mano != NINGUN_OBJETO 
     && se_puede_cortar( objeto_en_mano_personaje_activo( *juego ) ) ) {

        juego->comida[ juego->tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( *juego , stitch ) ].esta_cortado = true ;
    
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Verifica si el ingrediente se puede cocinar, según su tipo.
bool se_puede_cocinar ( char tipo_ingrediente ){
    bool se_puede = false ;
    switch ( tipo_ingrediente ) {
        case MASA :
            se_puede = true ;            
            break ;

        case CARNE :
            se_puede = true ;            
        break ;

        case MILANESA :
            se_puede = true ;            
        break ;

    }
    return se_puede ;
}



// PRE-CONDICIONES: El personaje activo debe ser reuben
//                  debe estar a una distancia manhattan de 1 con un horno y tener un ingrediente posible de cocinar
//POST-CONDICIONES: El ingrediente que tenga el mismo tipo que el que tiene reuben en la mano, cambiara su booleano "esta cocinado" a true 
//                  (Se cocina el ingrediente que tiene reuben en la mano)
void usar_horno ( juego_t *juego , personaje_t reuben ) {
    if( juego->personaje_activo == LETRA_REUBEN 
    && juego->reuben.objeto_en_mano != NINGUN_OBJETO
    && se_puede_cocinar ( objeto_en_mano_personaje_activo( *juego ) ) ) {

        for( int i=0 ; i<juego->tope_herramientas ; i++ ) {
            if ( ( calcular_distancia_manhattan ( reuben.posicion , juego->herramientas[i].posicion ) <= DISTANCIA_PARA_COCINAR )
                 && ( juego->herramientas[i].tipo == HORNOS)
                 && !juego->comida[juego->tope_comida-1].ingrediente[buscar_numero_ingrediente_en_mano ( *juego , reuben ) ].esta_cortado ) {

                     juego->comida[juego->tope_comida-1].ingrediente[buscar_numero_ingrediente_en_mano ( *juego , reuben ) ].esta_cocinado=true;
            }
        }
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Verifica si la mesa esta apta para dejar un ingrediente en ella.
bool condicion_mesa_vacia ( juego_t juego , personaje_t personaje_elegido ) {
    bool mesa_vacia = false ;
    if( ingrediente_en_mesa ( juego ) == NINGUN_OBJETO || ( ingrediente_en_mesa ( juego ) == personaje_elegido.objeto_en_mano ) ) mesa_vacia = true ;
    return mesa_vacia ;
}



// PRE-CONDICIONES: La mesa debe estar apta para dejar un ingrediente
//                  El personaje activo debe estar a una distancia manhattan de 1 y tener algún ingrediente.
//POST-CONDICIONES: La coordenada del ingrediente que tenga el mismo tipo que el del personaje cambiara a la coordenada de la mesa
//                  El objeto en mano del personaje cambiara a ningun objeto (deja el ingrediente en la mesa)
void dejar_ingrediente_mesa ( juego_t *juego , personaje_t *personaje ) {
    if (condicion_mesa_vacia ( *juego , *personaje ) ) {

        juego->comida[ juego->tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( *juego , *personaje ) ].posicion  =  juego->mesa ;
        personaje->objeto_en_mano = NINGUN_OBJETO ;
    }
}



// PRE-CONDICIONES: El personaje activo debe estar a una distancia manhattan de 1 y debe haber algún ingrediente en la mesa, 
//                  ese ingrediente no debe tener mismo tipo que el objeto en mano al de algún personaje
//POST-CONDICIONES: El objeto en mano del personaje activo cambia al tipo del ingrediente que esta en la mesa
void agarrar_ingrediente_mesa ( juego_t *juego , personaje_t *personaje ) {
    if ( juego->stitch.objeto_en_mano!=ingrediente_en_mesa( *juego ) && juego->reuben.objeto_en_mano!=ingrediente_en_mesa( *juego ) ) {

             personaje->objeto_en_mano = ingrediente_en_mesa( *juego ) ;
    }
}
// aclaracion**: cuando algun personaje agarra algo de la mesa, solo cambia su objeto en mano, el ingrediente sigue en la mesa y no se ve
// lo que da la posibilidad de que el otro personaje agarre el ingrediente de la mesa porque sigue ahí. Por eso los objeto en mano de los
// personajes deben ser distintos al ingrediente en la mesa.}



// PRE-CONDICIONES: -
//POST-CONDICIONES: El personaje activo usara la mesa, agarrara o dejara un ingrediente segun corresponda.
void usar_mesa ( juego_t *juego , personaje_t *personaje ) {
    if ( calcular_distancia_manhattan ( juego->mesa , personaje->posicion ) == DISTANCIA_PARA_USAR_MESA 
    && !hay_fuego_cuadrante_personaje_activo ( *juego , personaje->tipo )) {

        if ( personaje->objeto_en_mano == NINGUN_OBJETO  ) {
            agarrar_ingrediente_mesa ( &( *juego ) , &( *personaje ) ) ;

        }else if ( personaje->objeto_en_mano !=  MATAFUEGOS ) {
            dejar_ingrediente_mesa ( &( *juego ) , & ( *personaje ) ) ;
        }
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Mueve el ingredeinte elegido al final del vector y le resta 1 al tope
void borrar_ingrediente ( ingrediente_t *ingrediente_elegido , comida_t *comida_actual ) {
    ( *ingrediente_elegido ) = comida_actual->ingrediente[ comida_actual->tope_ingredientes - 1 ] ;
    (comida_actual->tope_ingredientes)-= 1 ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Verifica si el ingredeinte en mano del personaje indicado esta listo para emplatar.
bool ingrediente_en_mano_esta_listo ( juego_t juego , personaje_t personaje ) {
    bool esta_listo = false ;
    switch ( objeto_en_mano_personaje_activo ( juego  ) ) {
        case LECHUGA :
            if ( juego.comida[ juego.tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( juego , personaje ) ].esta_cortado) {
            esta_listo = true ;
            }
            break ;

        case TOMATE :
            if ( juego.comida[ juego.tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( juego , personaje ) ].esta_cortado ) {
            esta_listo = true ;
            }
        break ;

        case MASA :
            if ( juego.comida[ juego.tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( juego , personaje ) ].esta_cocinado ) {
            esta_listo = true ;
            }
        break ;

        case QUESO :
            if (juego.comida[ juego.tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( juego , personaje ) ].esta_cortado ) {
            esta_listo = true ;
            }
        break ;

        case JAMON :
            if ( juego.comida[ juego.tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( juego , personaje ) ].esta_cortado ) {
            esta_listo = true ;
            }
        break ;

        case PAN :
            if ( juego.comida[ juego.tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( juego , personaje ) ].esta_cortado ) {
            esta_listo = true ;
            }
        break ;

        case CARNE :
            if ( juego.comida[ juego.tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( juego , personaje ) ].esta_cocinado ) {
            esta_listo = true ;
            }
        break ;

        case MILANESA :
            if ( juego.comida[ juego.tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( juego , personaje ) ].esta_cocinado ) {
            esta_listo = true ;
            }
        break ;

        
    }
    return esta_listo ;
}



// PRE-CONDICIONES: El personaje indicado debe compartir coordenada con la salida (estar arriba de la salida)
//POST-CONDICIONES: El ingrediente que tiene el mismo tipo que el objeto en mano del personaje indicado se borra.  
void emplatar_ingrediente ( juego_t *juego , personaje_t *personaje ) {
    if ( objeto_debajo ( *juego , *personaje ) == SALIDA
    && ingrediente_en_mano_esta_listo( *juego , *personaje )
    && !hay_fuego_cuadrante_personaje_activo( *juego , juego->personaje_activo ) ) {

        juego->comida_lista[ juego->tope_comida_lista ]  =  juego->comida[ juego->tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano( (*juego) , juego->reuben ) ];
        juego->tope_comida_lista++ ;
        borrar_ingrediente ( &( juego->comida[ juego->tope_comida - 1 ].ingrediente[ buscar_numero_ingrediente_en_mano ( *juego , juego->reuben ) ] ) , 
                             &( juego->comida[ juego->tope_comida - 1 ] ) ) ;
        personaje->objeto_en_mano = NINGUN_OBJETO ;
        

    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Aparecera la siguiente comida cuando se cumplan las condiciones (cuando el plato anterior este completado)
void aparecer_siguiente_comida ( juego_t *juego){
    if (juego->tope_comida_lista == CANTIDAD_INGREDIENTES_ENSALADA 
        && juego->comida_actual == ENSALADA ) {

        inicializar_pizza( &(*juego) ) ;
        juego->tope_comida_lista = 0;
    }
    if (juego->tope_comida_lista == CANTIDAD_INGREDIENTES_PIZZA 
        && juego->comida_actual == PIZZA ) {
        inicializar_hamburguesa( &(*juego) ) ;
        juego->tope_comida_lista = 0;
    }
    if (juego->tope_comida_lista == CANTIDAD_INGREDIENTES_HAMBURGUESA 
        && juego->comida_actual == HAMBURGUESA) {
        inicializar_sandwich( &(*juego) ) ;
        juego->tope_comida_lista = 0;
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Encuentra la posicion del fuego, lo reemplaza por lo que este en el último lugar del vector y le baja el tope
//                  Encuentra la posicion del matafuegos, lo reemplaza por lo que este en el último lugar del vector y le baja el tope
void eliminar_fuego_matafuegos ( juego_t *juego ){
    for( int i=0 ; i<juego->tope_obstaculos ; i++ ){
        if( juego->obstaculos[i].tipo == FUEGO ){
                juego->obstaculos[i] =  juego->obstaculos[ juego->tope_obstaculos-1 ];
                juego->tope_obstaculos -=1 ;
        }
    }
    for( int i=0 ; i<juego->tope_herramientas ; i++ ){
        if( juego->herramientas[i].tipo == MATAFUEGOS ){
            juego->herramientas[i] =  juego->herramientas[ juego->tope_herramientas-1 ];
            juego->tope_herramientas -=1 ;
    }
}
    
    
}



// PRE-CONDICIONES: El personaje debe tener el matafuegis y estar a una distancia manhattan de 2
//POST-CONDICIONES: El fuego y el matafuegos se eliminan.
void usar_matafuegos ( juego_t *juego , personaje_t *personaje ) {
    if( calcular_distancia_manhattan ( personaje->posicion , juego->obstaculos[ juego->tope_obstaculos -1 ].posicion) <= DISTANCIA_PARA_USAR_MATAFUEGOS
        && personaje->objeto_en_mano == MATAFUEGOS) {
            eliminar_fuego_matafuegos( &(*juego) );
            personaje->objeto_en_mano = NINGUN_OBJETO ;
            juego->movimientos = 0 ;
        }
}



// PRE-CONDICIONES: El personaje debe tener misma coordenada que el matafuegos(estar arriba) y no tener ningun ingrediente en mano
//POST-CONDICIONES: El objeto en mano del perrsonaje cambia al matafuegos.
void agarrar_matafuegos( juego_t *juego , personaje_t *personaje ) {
    if(objeto_debajo( (*juego) , (*personaje)) == MATAFUEGOS && personaje->objeto_en_mano == NINGUN_OBJETO ) {
        personaje->objeto_en_mano = MATAFUEGOS;
    }
}



// PRE-CONDICIONES: No debe haber fuego y se deben haber acumulado 15 movimientos
//POST-CONDICIONES: Aparecera el fuego y el matafuego, asignandole una coordenada aleaotoria en un cuadrante aleaotorio
//                  sumandose en sus vectores correspondientes con sus convenciones correspondientes
void inicializar_fuego_matafuegos( juego_t *juego ) {
    if( !hay_fuego( *juego ) && juego->movimientos == 15){
        int cuadrante_aleatorio = numero_random ( 1 , 2 ) ;
        juego->obstaculos[ juego->tope_obstaculos ].posicion = generar_coordenada_aleatoria( cuadrante_aleatorio , (*juego) , juego->tope_comida - 1)  ;
        juego->obstaculos[ juego->tope_obstaculos ].tipo = FUEGO ;
        (juego->tope_obstaculos)++ ;
        juego->herramientas[ juego->tope_herramientas ].posicion = generar_coordenada_aleatoria ( cuadrante_aleatorio , (*juego) , juego->tope_comida -1 ) ;
        juego->herramientas[ juego->tope_herramientas ].tipo = MATAFUEGOS ;
        (juego->tope_herramientas)++ ;
    }
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Según el movimeinto dado, realizara la jugada correspondiente si es posible.
void realizar_jugada_personaje_elegido ( juego_t *juego , char movimiento , personaje_t *personaje_elegido  ) {
    switch ( movimiento ) {
        case LETRA_MOVER_ARRIBA :
            mover_personaje ( &(*juego) , &(*personaje_elegido) , VALOR_MOVIMIENTO_NEGATIVO , VALOR_MOVIMIENTO_NULO ) ; /*arriba*/
            break ;

        case LETRA_MOVER_ABAJO :
            mover_personaje (&(*juego) , &(*personaje_elegido) , VALOR_MOVIMIENTO_POSITIVO , VALOR_MOVIMIENTO_NULO ) ; /*abajo*/
            break;

        case LETRA_MOVER_IZQUIERDA:
            mover_personaje ( &(*juego) , &(*personaje_elegido) , VALOR_MOVIMIENTO_NULO , VALOR_MOVIMIENTO_NEGATIVO ) ; /*izquierda*/
            break;

        case LETRA_MOVER_DERECHA :
            mover_personaje (&(*juego) , &(*personaje_elegido) , VALOR_MOVIMIENTO_NULO , VALOR_MOVIMIENTO_POSITIVO ) ; /*derecha*/
            break ;

        case LETRA_USAR_MATAFUEGOS :
            usar_matafuegos( &(*juego) , &(*personaje_elegido) ) ;
            break ;

        case LETRA_CAMBIAR_PERSONAJE :
            cambiar_personaje ( &(*juego)) ;
            break ;

         case LETRA_AGARRAR_SOLTAR_INGREDIENTE :
            agarrar_soltar_ingrediente ( &(*juego) , &(*personaje_elegido) ) ;
            break ;    
    }

    if ( !hay_fuego_cuadrante_personaje_activo( *juego, personaje_elegido->tipo ) ){
        switch (movimiento) {
           

             case LETRA_USAR_MESA :
            usar_mesa ( &(*juego) , &(*personaje_elegido) ) ;
                 break ;

             case LETRA_USAR_CUCHILLO :
            usar_cuchillo ( &(*juego) , (*juego).stitch ) ;
                break ;

             case LETRA_USAR_HORNO :
            usar_horno ( &(*juego) , (*juego).reuben );
                break ;
        }
    }
    emplatar_ingrediente ( &(*juego) , &(*juego).reuben ) ;
    agarrar_matafuegos ( &(*juego) , &(*personaje_elegido) ) ;
}



// PRE-CONDICIONES: -
//POST-CONDICIONES: Reconoce el personaje activo y ejecuta la jugada con ese personaje
void ver_personaje_y_accionar_jugada ( juego_t *juego , char movimiento ) {
    switch ( juego->personaje_activo ) {
        case LETRA_STITCH :
            realizar_jugada_personaje_elegido ( &(*juego) , movimiento , &(*juego).stitch ) ;
            break ;

        case LETRA_REUBEN :
            realizar_jugada_personaje_elegido ( &(*juego) , movimiento , &(*juego).reuben ) ;
        break ;
    }
}






/* FUNCIONES PARA estado_juego*/




// PRE-CONDICIONES: -
//POST-CONDICIONES: Verifica si el juego esta perdido
bool condicion_juego_perdido ( juego_t juego ) {
    return ( objeto_debajo ( juego , juego.stitch ) == AGUJEROS || objeto_debajo( juego , juego.reuben ) == AGUJEROS ) ;
}


// PRE-CONDICIONES: -
//POST-CONDICIONES: Verifica si el juego esta ganado
bool condicion_juego_ganado ( juego_t juego ) {
    bool juego_ganado = false ;
    if ( juego.tope_comida_lista == 0 
        && juego.comida_actual == HAMBURGUESA
        && juego.precio_total <= PRECIO_ENSALADA_PIZZA_HAMBURGUESA ) {//<=100

            juego_ganado = true ;
    }

    if ( juego.tope_comida_lista == 0
        && juego.comida_actual == SANDWICH
        && juego.precio_total > PRECIO_ENSALADA_PIZZA_HAMBURGUESA
        && juego.precio_total <= PRECIO_ENSALADA_PIZZA_HAMBURGUESA_SANDWICH ) {// >100 Y <=150

            juego_ganado = true ;    
    }

    if ( juego.tope_comida_lista == CANTIDAD_INGREDIENTES_SANDWICH
        && juego.comida_actual == SANDWICH
        && juego.precio_total > PRECIO_ENSALADA_PIZZA_HAMBURGUESA_SANDWICH) {//>150

            juego_ganado = true ;
    }

    return juego_ganado;
}





/* PROCEDIMIENTOS Y FUNCIONES PRINCIPALES */




void inicializar_juego ( juego_t* juego , int precio ) {
    inicializar_topes ( &(*juego) ) ;
    inicializar_paredes ( (*juego).paredes , &(*juego).tope_paredes ) ;
    inicializar_agujeros ( &( *juego ) ) ;
    inicializar_herramientas ( &( *juego ) ) ;
    inicializar_ensalada( &(*juego ) ) ;
    inicializar_personajes ( &( *juego ) ) ;
    juego->salida = generar_coordenada_aleatoria ( CUADRANTE_REUBEN , (*juego) , juego->tope_comida -1 ) ;
    inicializar_mesa( &( *juego ) ) ;
    juego->precio_total = precio ;
    juego->movimientos = 0 ;
}


void realizar_jugada ( juego_t *juego , char movimiento ) {
    ver_personaje_y_accionar_jugada ( &(*juego) , movimiento ) ;
    aparecer_siguiente_comida ( &( *juego ) ) ;
    inicializar_fuego_matafuegos ( &(*juego) ) ;
}


void imprimir_terreno(juego_t juego){
    char matriz[MAXIMO_MATRIZ][MAXIMO_MATRIZ] ;
    cargar_matriz ( matriz ) ;
    cargar_paredes ( juego , matriz ) ;
    cargar_obstaculos ( juego , matriz ) ;
    cargar_herramientas ( juego , matriz ) ;
    cargar_mesa_y_salida ( juego , matriz ) ;
    cargar_ingredientes ( juego , matriz ) ;
    cargar_personajes(juego, matriz) ;
     printf(" ██████╗ ██╗   ██╗███████╗██████╗        ██████╗██╗  ██╗ █████╗ ███╗   ███╗██████╗ ██╗   ██╗ ██████╗██╗  ██╗██╗████████╗ ██████╗ \n"
           "██╔═══██╗██║   ██║██╔════╝██╔══██╗      ██╔════╝██║  ██║██╔══██╗████╗ ████║██╔══██╗██║   ██║██╔════╝██║  ██║██║╚══██╔══╝██╔═══██╗\n"
           "██║   ██║██║   ██║█████╗  ██████╔╝█████╗██║     ███████║███████║██╔████╔██║██████╔╝██║   ██║██║     ███████║██║   ██║   ██║   ██║\n"
           "██║   ██║╚██╗ ██╔╝██╔══╝  ██╔══██╗╚════╝██║     ██╔══██║██╔══██║██║╚██╔╝██║██╔══██╗██║   ██║██║     ██╔══██║██║   ██║   ██║   ██║\n"
           "╚██████╔╝ ╚████╔╝ ███████╗██║  ██║      ╚██████╗██║  ██║██║  ██║██║ ╚═╝ ██║██████╔╝╚██████╔╝╚██████╗██║  ██║██║   ██║   ╚██████╔╝\n"
           " ╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═╝       ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚═════╝  ╚═════╝  ╚═════╝╚═╝  ╚═╝╚═╝   ╚═╝    ╚═════╝ \n");
    mostrar_matriz(matriz) ;
    if (juego.personaje_activo == LETRA_STITCH){
            mostrar_info_personaje( juego, juego.stitch);
        } else mostrar_info_personaje( juego, juego.reuben) ;
    
    printf("cantidad movimientos %i\n", juego.movimientos) ; 
}


int estado_juego ( juego_t juego ) {
    int resultado_juego = SIGUE_JUEGO ;
    if ( condicion_juego_perdido ( juego ) ) {
        resultado_juego = JUEGO_PERDIDO ;
    }
    if ( condicion_juego_ganado ( juego ) ) {
        resultado_juego = JUEGO_GANADO ;
    }
    return resultado_juego ;
}


