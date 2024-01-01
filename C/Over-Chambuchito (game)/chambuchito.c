#include <stdio.h>
#include "chambuchito.h"
const int TAMANO_MINIMO=15;
const int TAMANO_MAXIMO=30;
const char PAN_BLANCO= 'B';
const char PAN_INTEGRAL= 'I';
const char PAN_AVENA_MIEL= 'A';
const char PAN_QUESO_OREGANO= 'Q';
const char QUESO_DAMBO= 'D';
const char QUESO_CHEDDAR= 'C';
const char QUESO_GRUYERE= 'G';
const char SIN_QUESO= 'S';
const char ROAST_BEEF= 'R';
const char ATUN= 'A';
const char SOJA= 'S';
const char POLLITO= 'P';
const char NADA_DE_PROTE= 'N';
const char RESPUESTA_SI= 'S';
const char RESPUESTA_NO= 'N';
const double VALOR_CALCULO_PRECIO_MEDIDA= 0.3;
const int PRECIO_PAN_BASICO= 5;
const int PRECIO_PAN_ESPECIAL= 8;
const int PRECIO_QUESO_BASICO= 5;
const int PRECIO_QUESO_ESPECIAL= 8;
const int PRECIO_ROAST_BEEF=7;
const int PRECIO_ATUN=9;
const int PRECIO_POLLITO=5;
const int PRECIO_SOJA=3;

//Pre:-
//Post: Pide y guarda la medida del sanguchito elegida por el usuario.
//      Esa medida estara entre 15 y 30.
void pregunta_tamano(int* tamano_elegido){
    while(*tamano_elegido<TAMANO_MINIMO || *tamano_elegido>TAMANO_MAXIMO){
        printf("\nDe que medida desea su chambuchito? Debe ser entre 15cm y 30cm.\nSolamente ingrese el numero, sin el -cm-\n");
        scanf("%i", &*tamano_elegido);
    }}

//Pre:-
//Post:Pide y guarda el tipo de pan elegido por el usuario.
//     Se repetira la pregunta hasta que se elija una opci�n.
void pregunta_tipo_pan(char* pan_elegido){
    while(!(*pan_elegido==PAN_BLANCO || *pan_elegido==PAN_INTEGRAL || *pan_elegido==PAN_AVENA_MIEL || *pan_elegido==PAN_QUESO_OREGANO)){
        printf("\nElija que pancito desea en su chambuchito\nBlanco -B-, Integral -I-, Avena y Miel -A-, Queso y Oregano -Q-\nIngrese solo la letra correspondiente de su pancito, en mayuscula\n");
        scanf( " %c", &*pan_elegido);
    }}  

//pre-
//Post: Pide y guarda el tipo de queso elegido por el usuario.
//     Se repetira la pregunta hasta que se elija una opci�n.
void pregunta_tipo_queso(char* queso_elegido){
    while(!(*queso_elegido==QUESO_DAMBO || *queso_elegido==QUESO_CHEDDAR || *queso_elegido==QUESO_GRUYERE || *queso_elegido==SIN_QUESO)){
        printf("\nElija que queso quiere en su chambuchito\nDambo -D-, Cheddar -C-, Gruyere -G-, Sin Queso -S-.\nIngrese solo la letra correspondiente de su quesito, en mayuscula\n");
        scanf( " %c", &*queso_elegido);
    }}  

//pre:-
//post: Pide y guarda el tipo de prote elegido por el usuario.
//      Se repetira la pregunta hasta que se elija una opci�n.
void pregunta_tipo_prote(char* prote_elegida){
    while(!(*prote_elegida==ROAST_BEEF || *prote_elegida==ATUN || *prote_elegida==POLLITO || *prote_elegida==SOJA || *prote_elegida==NADA_DE_PROTE)){
        printf("\nElija que prote quiere en su chambuchito\nRoast Beef -R-, ATUN -A-, SOJA -S-, POLLITO -P-, Nada de prote -N-\nIngrese solo la letra correspondiente de su prote, en mayuscula\n");
        scanf( " %c", &*prote_elegida);
    }}  

//pre: La prote no debe ser atun
//post:Pregunta si se debe calentar o no el sanguchito.
void pregunta_calentar(char* prote_elegida , char* respuesta_calentar){
    if(!(*prote_elegida==ATUN)){
        while(!(*respuesta_calentar==RESPUESTA_SI || *respuesta_calentar==RESPUESTA_NO)){
            printf("\n�Quiere calentar su chambichito?\nSi -S-  ,   No -N-\nIngrese solo la letra correspondiente en mayuscula\n");
            scanf( " %c", &*respuesta_calentar);
     }}}    

//pre:se debe haber elegido la medida del sanguchito
//post: calcula y asigna el precio de la medida del sanguchito
void asignacion_precio_medida(int tamano_elegido, double *precio_medida){
    double tamano_elegido_double=(double)tamano_elegido;
    *precio_medida=VALOR_CALCULO_PRECIO_MEDIDA*tamano_elegido_double;
}

//pre:se debe haber elegido el pan
//post:calcula y asigna el precio del pancito del sanguchito
void asignacion_precio_pan( char pan_elegido, int* precio_pan){
    if( pan_elegido==PAN_BLANCO || pan_elegido==PAN_INTEGRAL){
        *precio_pan=PRECIO_PAN_BASICO;
    } else {
        *precio_pan=PRECIO_PAN_ESPECIAL;
    }
}

//pre:se debe haber elegido el queso
//post:calcula y asigna el precio del quesito del sanguchito
void asignacion_precio_queso( char queso_elegido, int* precio_queso){
    if( queso_elegido==QUESO_DAMBO || queso_elegido==QUESO_CHEDDAR){
        *precio_queso=PRECIO_QUESO_BASICO;
    } else if(queso_elegido==QUESO_GRUYERE){
        *precio_queso=PRECIO_QUESO_ESPECIAL;
    }
}

//pre:se debe haber elegido la prote
//post:calcula y asigna el precio de la prote del sanguchito
void asignacion_precio_prote(char prote_elegida, int* precio_prote ){
    if( prote_elegida==ROAST_BEEF){
        *precio_prote=PRECIO_ROAST_BEEF;
    } else if(prote_elegida==ATUN){
        *precio_prote=PRECIO_ATUN;
    } else if(prote_elegida==POLLITO){
        *precio_prote=PRECIO_POLLITO;
    } else if(prote_elegida==SOJA){
        *precio_prote=PRECIO_SOJA;
    }
}

//pre: los precios de cada ingrediente y la medida deben estar asignados
//post:calcula el precio final del sanguchito sin redondear
double calculo_precio_final( double precio_medida, int precio_pan, int precio_queso, int precio_prote, double *precio_final){

    *precio_final= ( (double)(precio_pan+precio_prote+precio_queso))*(precio_medida);
    return *precio_final;
}

//pre:se debe haber calculado el precio final redondeado del sanguchito
//post:muestra el precio final redondeado del sanguchito
void mostrar_precio(int precio_final_entero){
    printf("Su precio final es -%i-\n",precio_final_entero);
}

void calcular_precio_chambuchito(int *precio){
    int tamano_elegido=0;
    char pan_elegido;
    char queso_elegido;
    char prote_elegida;
    char respuesta_calentar;
    double precio_medida=0;
    int precio_pan=0;
    int precio_queso=0;
    int precio_prote=0;
    pregunta_tamano( &tamano_elegido);
    pregunta_tipo_pan( &pan_elegido);
    pregunta_tipo_queso( &queso_elegido);
    pregunta_tipo_prote( &prote_elegida);
    pregunta_calentar( &prote_elegida, &respuesta_calentar);
    asignacion_precio_medida(tamano_elegido, &precio_medida);
    asignacion_precio_pan(pan_elegido, &precio_pan);
    asignacion_precio_queso( queso_elegido, &precio_queso);
    asignacion_precio_prote( prote_elegida, &precio_prote);
    double precio_final=calculo_precio_final(precio_medida, precio_pan, precio_queso, precio_prote, &precio_final);
    int precio_final_entero=(int)precio_final;
    mostrar_precio (precio_final_entero);
    *precio=precio_final_entero;
}
