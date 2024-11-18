La situación actual de la implementación es la siguiente:

Se crea una interfaz registros donde se tiene:
    registroActual string: es una cadena en la que se guarda qué registro se está analizando en ese momento (para detectar bien los ataques DoS)
    diccionarioIPs: es un ABB para guardar las IPs como clave y una estructura datos_diccionario como dato con la siguiente estructura...
        type datos_diccionario struct {
	        ultimaVisita       string //se guarda en qué registro se realizó la última visita a esta IP
	        tiempo             time.Time //se guarda el tiempo en el cual se registró la primera entrada (para hacer las verificaciones de DoS)
	        visitasDesdeTiempo int //es un contador de la cantidad de visitas que se realizaron desde tiempo (para hacer las verificaciones de DoS)
            ataqueDoSReportado bool //bool indicando si ya se reportó un ataque DoS para esa IP
        }
    un diccionarioSitiosVisitados que contiene como
            clave = sitio  string           //nombre del sitio
            dato = cantidad_visitas int    //contador de la cantidad de veces que se visitó el sitio
        El objetivo es guardar los sitios en el hash para poder actualizar en O(1).
        Para hallar los n sitios más visitados se itera el hash y se van guardando los sitiosVisitados en un heap de máximos ordenados por su cantidad de visitas.
        Al desencolar del heap de máximos salen los sitios visitados en orden descendente. 

Adicionalmente se agrega en el TDARegistros una función RealizarOperacion() que realiza la operación correspondiente, indicando con un booleano si la pudo realizar correctamente.