package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {

	tlv := []byte{49, 49, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48}
	//tlv := []byte("11A05AB398765UJ102N2300")
	mapeo := tlv_decode(tlv)

	println(mapeo["value"])
}

func tlv_decode(tlv_byte_code []byte) map[string]string {

	final_map := map[string]string{}
	tlv := string(tlv_byte_code)
	data := map[string]string{}
	//cont := 1
	for ok := true; ok; ok = (len(tlv) > 0) {
		if len(tlv) > 0 {

			largo, err := strconv.Atoi(tlv[0:2])
			if err == nil {
				//fmt.Println(largo)
			}

			//tipo := tlv[2:5]
			//fmt.Println(tipo)

			valor := tlv[5 : 5+largo]
			//fmt.Println(valor)
			//data["id"] = fmt.Sprintf("%b", cont)
			data["valor"] = valor
			data["error"] = "200"
			tlv = tlv[5+largo:]
			mapFinal, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				fmt.Println("error:", err)
			}
			print(string(mapFinal))

		} else {
			data["valor"] = "Empty"
			data["error"] = "Cadena vacia"
			//fmt.Println("Cadena vacia")

		}
	}

	return final_map

}
