#include <stdio.h>
#include "chambuchito.h"
#include "cocineritos.h"
#include <time.h> 
#include <stdlib.h> 

const static int SIGUE_JUEGO = 0;

int main(){
    srand ((unsigned)time(NULL));    
    juego_t juego;    
    int precio = 300 ;
    char movimiento =' ';
    int resultado_juego = SIGUE_JUEGO ;
   // calcular_precio_chambuchito( &precio ) ;
    inicializar_juego( &juego , precio ) ;
    while ( resultado_juego == SIGUE_JUEGO ) {
        system( "clear" ) ; 
        imprimir_terreno( juego ) ;
        scanf( " %c", &movimiento ) ;
        realizar_jugada (&juego , movimiento ) ;
        resultado_juego = estado_juego( juego ) ;
    }
    return 0;
}