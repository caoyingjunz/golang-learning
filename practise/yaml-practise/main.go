package main

import gopkg.in/yaml.v2

//refer to https://github.com/go-yaml/yaml
// https://www.jianshu.com/p/84499381a7da

func main() {
	// resultMap := make(map[string]interface{})
	conf := new(module.Yaml)
	yamlFile, err := ioutil.ReadFile("test.yaml")

	// conf := new(module.Yaml1)
	// yamlFile, err := ioutil.ReadFile("test.yaml")

	// conf := new(module.Yaml2)
	//  yamlFile, err := ioutil.ReadFile("test1.yaml")

	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
	// log.Println("conf", resultMap)
}
