## Testing

### Code coverage

Parado en la raíz del proyecto a ejecutar

```cli
go test -cover ./...
```
Gnerear el coverage report, ejecutar:
```cli
go test -cover -coverprofile=coverage.out  ./...
```
Para ver el porcentaje de cobertura por función
```bash
go tool cover -func coverage.out
````
Mostrar en el navegador el coverage report generado, ejecutar:

```cli
go tool cover -html=coverage.out
```



* ¿Cuáles son las diferencias entre White Box y Black Box? 
  
  En el black box solo se conoce las entradas y salidas del sistema, se desconoce por completo el cómo, solo el qué. Mientras que en el white box se conoce el funcionamiento interno (el ódigo es visible), esto hace el proceso de testing más complicado ya que se debe contemplar las casuísticas de cada unidad y las relaciones entre unidades (tests unitarios y tests de integración).

* ¿Qué es un test funcional?

En los tests funcionales se evalúa que la aplicación cumpla con los requerimientos definidos. Se aplica la metodoloía de black box ya que solo se evalúa que la salida sea acorde con los requrimientos y entradas proporcionadas

* ¿Qué es un Test de Integración?

En el test de integración se evalúa que las capas o unidades se comuniquen bien entre sí y que el ecosistema como un todo funcione bien

* Indicar las dimensiones de calidad prioritarias en MELI.
  
Meli se basa en el estándar de calidad ISO/IEC 25010  y prioriza cinco de sus principios:

* Funcionalidad: tiene que ver con el cumplimiento de los requerimientos definidos
  
* Rendimiento: tiene que ver con el manejo de recursos (tiempo de ejecución y uso de memoria).
  
* Fiabilidad: tiene que ver con la confiablidad del código, se espera que sea poco propenso a errores y que sea pueda recuperar fácilmente en caso de que ocurran. Se podría medir en función de un escenario y tiempo definido.
  
* Seguridad: evitar puntos vulnerables aplicando programación defensiva en la que se tratan de considerar la mayor cantidad de escenarios posibles.
  
* Mantenibilidad: tiene que ver con baja cohesión entre las capas(inversión de dependencias), reusabilidad y testeabilidad. El código debería ser fácil de mantener por otros miembros del equipo e inclusive personas ajenas al proyecto.
  
