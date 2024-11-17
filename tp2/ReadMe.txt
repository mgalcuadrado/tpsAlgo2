La situación actual de la implementación es esta:

Se crea en un diccionario (hash) de funcionesDisponibles donde:
    clave = comando enviado por línea de comandos argv[0]
    dato = cantidad de argumentos esperados para ese argumento argc
    Esto tiene el objetivo de permitirnos verificar fácilmente que se pueda realizar la función
    Me gustaría encontrar una forma de incluirle la función como dato y poder directamente llamarla así

Se crea una interfaz registros donde se tiene:
    un ABB para guardar las IPs como clave y una estructura datos_diccionario como dato donde...
        type datos_diccionario struct {
	        ultimaVisita       log //se guarda en qué registro se realizó la última visita a esta IP
	        tiempo             time.Time //se guarda el tiempo en el cual se registró la primera entrada (para hacer las verificaciones de DoS)
	        visitasDesdeTiempo int //es un contador de la cantidad de visitas que se realizaron desde tiempo (para hacer las verificaciones de DoS)
            ataqueDoSReportado bool //bool indicando si ya se reportó un ataque DoS para esa IP
        }
    un hashSitiosVisitados contiene como
            clave = sitio            string //nombre del sitio
            dato = cantidad_visitas int    //contador de la cantidad de veces que se visitó el sitio
        El objetivo es guardar los sitios en el hash para poder actualizar en O(1), y después pasarlo a un arreglo y hacer HeapSort para sacarlos ordenados por máximos