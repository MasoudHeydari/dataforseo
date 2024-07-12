package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main01() {
	// Define the request body
	requestBody := []map[string]interface{}{
		{
			"keyword":       "expanded polystyrene",
			"location_code": 1003854,
			"language_code": "en",
			"depth":         100,
			//"device":        "desktop",
			//"os":            "windows",
		},
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error encoding request body:", err)
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest(http.MethodPost, "https://api.dataforseo.com/v3/serp/google/organic/task_post", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set the headers
	token := base64.StdEncoding.EncodeToString([]byte("m.heydari4883@gmail.com:22599da38215faea"))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()
	respStr, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io.ReadAll: ", err)
	}

	// Print the response status code
	fmt.Println("Response status code:", resp.StatusCode)
	fmt.Println(respStr)
}

func main() {
	// Create the HTTP request
	u := "https://api.dataforseo.com/v3/serp/google/organic/task_get/regular/07111532-8040-0066-0000-4bf39ad568ad"
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set the headers
	token := base64.StdEncoding.EncodeToString([]byte("m.heydari4883@gmail.com:22599da38215faea"))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()
	//respStr, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("io.ReadAll: ", err)
	//	return
	//}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	//response := new(Response)
	//err = json.Unmarshal(respStr, resp)
	//if err != nil {
	//	// handle error
	//	fmt.Println("json.Unmarshal: ", err)
	//	return
	//}
	fmt.Println(response)
}

/*
{
    "version": "0.1.20240626",
    "status_code": 20000,
    "status_message": "Ok.",
    "time": "0.0964 sec.",
    "cost": 0,
    "tasks_count": 1,
    "tasks_error": 0,
    "tasks": [
        {
            "id": "07111532-8040-0066-0000-4bf39ad568ad",
            "status_code": 20000,
            "status_message": "Ok.",
            "time": "0.0298 sec.",
            "cost": 0,
            "result_count": 1,
            "path": [
                "v3",
                "serp",
                "google",
                "organic",
                "task_get",
                "regular",
                "07111532-8040-0066-0000-4bf39ad568ad"
            ],
            "data": {
                "api": "serp",
                "function": "task_get",
                "se": "google",
                "se_type": "organic",
                "depth": 100,
                "keyword": "expanded polystyrene",
                "language_code": "en",
                "location_code": 1003854,
                "device": "desktop",
                "os": "windows"
            },
            "result": [
                {
                    "keyword": "expanded polystyrene",
                    "type": "organic",
                    "se_domain": "google.de",
                    "location_code": 1003854,
                    "language_code": "en",
                    "check_url": "https:\/\/www.google.de\/search?q=expanded%20polystyrene&num=100&hl=en&gl=DE&gws_rd=cr&ie=UTF-8&oe=UTF-8&glp=1&uule=w+CAIQIFISCQFZAz83TqhHEXA7XltGICEE",
                    "datetime": "2024-07-11 12:32:39 +00:00",
                    "spell": null,
                    "item_types": [
                        "featured_snippet",
                        "people_also_ask",
                        "organic",
                        "popular_products",
                        "related_searches"
                    ],
                    "se_results_count": 11600000,
                    "items_count": 100,
                    "items": [
                        {
                            "type": "featured_snippet",
                            "rank_group": 1,
                            "rank_absolute": 1,
                            "domain": "www.bpf.co.uk",
                            "title": "Expanded Polystyrene (EPS)",
                            "description": "WHAT IS EPS? Expanded Polystyrene (EPS) is a rigid, closed cell, thermoplastic foam material produced from solid beads of polystyrene, which is polymerised from styrene monomer and contains an expansion gas (pentane) dissolved within the polystyrene bead.",
                            "url": "https:\/\/www.bpf.co.uk\/plastipedia\/polymers\/expanded-and-extruded-polystyrene-eps-xps.aspx",
                            "breadcrumb": null
                        },
                        {
                            "type": "organic",
                            "rank_group": 1,
                            "rank_absolute": 3,
                            "domain": "en.wikipedia.org",
                            "title": "Polystyrene",
                            "description": "Polystyrene is flammable, and releases large amounts of black smoke upon burning. · Expanded polystyrene is lightweight. This is a man in Guiyang, China carrying ...",
                            "url": "https:\/\/en.wikipedia.org\/wiki\/Polystyrene",
                            "breadcrumb": "https:\/\/en.wikipedia.org › wiki › Polystyrene"
                        },
                        {
                            "type": "organic",
                            "rank_group": 2,
                            "rank_absolute": 5,
                            "domain": "insulationcorp.com",
                            "title": "expanded polystyrene (EPS) foam",
                            "description": "Expanded Polystyrene (EPS) foam is a lightweight, rigid, closed-cell insulation available in various densities to withstand load and back-fill forces.",
                            "url": "https:\/\/insulationcorp.com\/eps\/",
                            "breadcrumb": "https:\/\/insulationcorp.com › eps"
                        },
                        {
                            "type": "organic",
                            "rank_group": 3,
                            "rank_absolute": 6,
                            "domain": "plasticseurope.org",
                            "title": "Expanded polystyrene (EPS)",
                            "description": "Expanded polystyrene, or EPS, is a widely used commodity polymer. It has been a material of choice for more than 50 years because of its versatility, ...",
                            "url": "https:\/\/plasticseurope.org\/plastics-explained\/a-large-family\/expanded-polystyrene\/",
                            "breadcrumb": "https:\/\/plasticseurope.org › ... › A large family"
                        },
                        {
                            "type": "organic",
                            "rank_group": 4,
                            "rank_absolute": 7,
                            "domain": "omnexus.specialchem.com",
                            "title": "Expanded Polystyrene (EPS Foam): Uses, Structure & ...",
                            "description": "Expanded polystyrene (EPS) is a lightweight and rigid foam material. It is a material of choice for the packaging and construction industry.",
                            "url": "https:\/\/omnexus.specialchem.com\/selection-guide\/expanded-polystyrene-eps-foam-insulation",
                            "breadcrumb": "https:\/\/omnexus.specialchem.com › selection-guide › exp..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 5,
                            "rank_absolute": 8,
                            "domain": "www.knauf-industries.com",
                            "title": "Know all about Expanded Polystyrene EPS",
                            "description": "Expandable polystyrene, our raw material, is manufactured in part by Knauf in Belgium. It is obtained by polymerising styrene and an expansion agent: pentane.",
                            "url": "https:\/\/www.knauf-industries.com\/en\/expanded-polystyrene\/",
                            "breadcrumb": "https:\/\/www.knauf-industries.com › expanded-polystyre..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 6,
                            "rank_absolute": 9,
                            "domain": "www.amazon.de",
                            "title": "Neopor® NEO WLG 032 80 mm Expanded Polystyrene ...",
                            "description": "Neopor® NEO WLG 032 80 mm Expanded Polystyrene (EPS) External Wall Insulation Boards, WDVS Styrofoam, 3 m² : Amazon.de: DIY & Tools.",
                            "url": "https:\/\/www.amazon.de\/-\/en\/Expanded-Polystyrene-External-Insulation-Styrofoam\/dp\/B07PJRK4NW",
                            "breadcrumb": "https:\/\/www.amazon.de › Expanded-Polystyrene-Extern..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 7,
                            "rank_absolute": 10,
                            "domain": "www.sulzer.com",
                            "title": "Expandable polystyrene (EPS)",
                            "description": "Expandable polystyrene (EPS) consists of polystyrene micro-pellets or beads containing a blowing agent and other additives for foaming.",
                            "url": "https:\/\/www.sulzer.com\/en\/shared\/products\/expandable-polystyrene-eps",
                            "breadcrumb": "https:\/\/www.sulzer.com › shared › products › expandabl..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 8,
                            "rank_absolute": 11,
                            "domain": "ceresana.com",
                            "title": "Expandable Polystyrene (EPS) Market Report",
                            "description": "According to the latest, and already fifth, edition of the “Expanded Polystyrene” market report, a total of around 7.2 million tonnes of EPS were consumed ...",
                            "url": "https:\/\/ceresana.com\/en\/produkt\/expandable-polystyrene-market-report",
                            "breadcrumb": "https:\/\/ceresana.com › ... › Products › Plastics"
                        },
                        {
                            "type": "organic",
                            "rank_group": 9,
                            "rank_absolute": 12,
                            "domain": "www.coperion.com",
                            "title": "Single-stage Production of Expandable Polystyrene (EPS)",
                            "description": "Expandable polystyrene (EPS) is a rigid and tough foam. EPS is often used for food packaging, building insulation, and packing material either as solid ...",
                            "url": "https:\/\/www.coperion.com\/en\/industries\/plastics\/expandable-polystyrene-eps",
                            "breadcrumb": "https:\/\/www.coperion.com › industries › plastics › expa..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 10,
                            "rank_absolute": 13,
                            "domain": "www.styropor.com",
                            "title": "BASF",
                            "description": "Styropor ® an expandable polystyrene (EPS), was invented in 1951 by BASF and has become a classic among raw materials in the construction and packaging industry ...",
                            "url": "https:\/\/www.styropor.com\/portal\/basf\/en\/dt.jsp?page=en",
                            "breadcrumb": "https:\/\/www.styropor.com › portal › basf"
                        },
                        {
                            "type": "organic",
                            "rank_group": 11,
                            "rank_absolute": 14,
                            "domain": "www.amazon.de",
                            "title": "Green Bean © Expanded Polystyrene (EPS) Beads ...",
                            "description": "Green Bean © Expanded Polystyrene (EPS) Beads, Highest Premium Quality, Refill Pack, Bean Bag Filling, Polystyrene Balls : Amazon.de: Home & Kitchen.",
                            "url": "https:\/\/www.amazon.de\/-\/en\/Expanded-Polystyrene-Highest-Premium-Quality\/dp\/B08CXTDDR8",
                            "breadcrumb": "https:\/\/www.amazon.de › Expanded-Polystyrene-Highe..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 12,
                            "rank_absolute": 15,
                            "domain": "knaufappliances.com",
                            "title": "Expanded Polystyrene - EPS",
                            "description": "Thermal and Acoustic Insulation. Due to the predominance of air in its composition, expanded polystyrene is an excellent thermal and acoustic insulator. This ...",
                            "url": "https:\/\/knaufappliances.com\/expanded-polystyrene-eps\/",
                            "breadcrumb": "https:\/\/knaufappliances.com › expanded-polystyrene-eps"
                        },
                        {
                            "type": "organic",
                            "rank_group": 13,
                            "rank_absolute": 16,
                            "domain": "www.panasorb.eu",
                            "title": "Styropor EPS polystyrene hard foam material sheet 100 x ...",
                            "description": "Order polystyrene from stock online. Large selection and rapid delivery from stock. Order your polystyrene directly from us.",
                            "url": "https:\/\/www.panasorb.eu\/lng\/en\/polystyrene\/styropor-eps-polystyrene-hard-foam-material-sheet-100-x-50-x-2cm.html?language=en",
                            "breadcrumb": "https:\/\/www.panasorb.eu › Polystyrene"
                        },
                        {
                            "type": "organic",
                            "rank_group": 14,
                            "rank_absolute": 17,
                            "domain": "www.engineeredfoamproducts.com",
                            "title": "Expanded Polystyrene | EPS Foam Manufacturer",
                            "description": "Expanded polystyrene can expand up to 50 times its original size. Learn more about EPS foam and its multiple properties from leading manufacturers.",
                            "url": "https:\/\/www.engineeredfoamproducts.com\/materials\/expanded-polystyrene\/",
                            "breadcrumb": "https:\/\/www.engineeredfoamproducts.com › Materials"
                        },
                        {
                            "type": "organic",
                            "rank_group": 15,
                            "rank_absolute": 18,
                            "domain": "www.bosig.de",
                            "title": "Expanded polystyrene foam (EPS)",
                            "description": "Expanded polystyrene foam (EPS) is a closed-cell insulation material that is manufactured from expanded polystyrol granulate and foamed up into blocks. EPS is ...",
                            "url": "https:\/\/www.bosig.de\/en\/produkte\/expanded-polystyrene-foam-eps\/",
                            "breadcrumb": "https:\/\/www.bosig.de › produkte › expanded-polystyrene..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 16,
                            "rank_absolute": 19,
                            "domain": "europlas.com.vn",
                            "title": "Expanded Polystyrene (EPS): What is it? - EuroPlas",
                            "description": "Expanded Polystyrene (EPS) is a type of rigid foam plastic that is derived from polystyrene beads. It is created through a process called expansion, where ...",
                            "url": "https:\/\/europlas.com.vn\/en-US\/blog-1\/expanded-polystyrene-eps-what-is-it",
                            "breadcrumb": "https:\/\/europlas.com.vn › Home › Blog"
                        },
                        {
                            "type": "organic",
                            "rank_group": 17,
                            "rank_absolute": 20,
                            "domain": "epsa.org.au",
                            "title": "What is EPS?",
                            "description": "EPS is used widely in the building and construction industry. EPS is an inert material that does not rot and provides no nutritional benefits to vermin ...",
                            "url": "https:\/\/epsa.org.au\/what-is-eps\/",
                            "breadcrumb": "https:\/\/epsa.org.au › what-is-eps"
                        },
                        {
                            "type": "organic",
                            "rank_group": 18,
                            "rank_absolute": 21,
                            "domain": "de.wikipedia.org",
                            "title": "Polystyrol",
                            "description": "Expandiertes Polystyrol (EPS, vor allem bekannt unter dem Handelsnamen Styropor) und extrudiertes Polystyrol (XPS) werden als Schaumstoffe eingesetzt.",
                            "url": "https:\/\/de.wikipedia.org\/wiki\/Polystyrol",
                            "breadcrumb": "https:\/\/de.wikipedia.org › wiki › P..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 19,
                            "rank_absolute": 22,
                            "domain": "www.gap-polymers.com",
                            "title": "Expanded vs Extruded Polystyrene: How are they different?",
                            "description": "Regarding the production process, EPS is manufactured by extruding a thick, hot fluid of PS crystals. Therefore, the critical difference is that ...",
                            "url": "https:\/\/www.gap-polymers.com\/en\/blog-post\/expanded-vs-extruded-polystyrene",
                            "breadcrumb": "https:\/\/www.gap-polymers.com › blog-post › expanded-..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 20,
                            "rank_absolute": 23,
                            "domain": "circulareconomy.europa.eu",
                            "title": "European Manufacturers of Expanded Polystyrene (EUMEPS)",
                            "description": "Plastic materials: ... Expanded Polystyrene is a particle foam. Its protective and lightweight properties make it particularly suitable for both insulation of ...",
                            "url": "https:\/\/circulareconomy.europa.eu\/platform\/en\/commitments\/pledges\/european-manufacturers-expanded-polystyrene-eumeps",
                            "breadcrumb": "https:\/\/circulareconomy.europa.eu › platform › pledges"
                        },
                        {
                            "type": "organic",
                            "rank_group": 21,
                            "rank_absolute": 24,
                            "domain": "www.thoughtco.com",
                            "title": "What Is EPS or Expanded Polystyrene?",
                            "description": "EPS is inert in nature and therefore does not result in any chemical reactions. Since it will not appeal to any pests, it can be used easily in ...",
                            "url": "https:\/\/www.thoughtco.com\/what-is-eps-expanded-polystyrene-820450",
                            "breadcrumb": "https:\/\/www.thoughtco.com › Chemistry › Basics"
                        },
                        {
                            "type": "organic",
                            "rank_group": 22,
                            "rank_absolute": 25,
                            "domain": "www.sensxpert.com",
                            "title": "Expanded Polystyrene (EPS) | sensXPERT Glossary",
                            "description": "Expanded polystyrene (EPS) is a popular thermoplastic foam valued for its combination of lightweight properties, thermal insulation, and affordability.",
                            "url": "https:\/\/www.sensxpert.com\/glossary\/expanded-polystyrene-eps\/",
                            "breadcrumb": "https:\/\/www.sensxpert.com › glossary › expanded-polyst..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 23,
                            "rank_absolute": 26,
                            "domain": "www.bioplaster-research.com",
                            "title": "Expanded Polystyrene, friend or foe? - BioPlaster Research",
                            "description": "Its characteristics offer us endless possibilities; however, EPS is a double-edged sword since this material has a significant environmental impact due to its ...",
                            "url": "https:\/\/www.bioplaster-research.com\/blog-english\/expanded-polystyrene-friend-or-foe",
                            "breadcrumb": "https:\/\/www.bioplaster-research.com › blog-english › e..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 24,
                            "rank_absolute": 27,
                            "domain": "www.novapor.com",
                            "title": "EPS - Expanded polystyrene - Styrofoam",
                            "description": "TYPICAL PROPERTIES OF STYROFOAM (EPS). Due to its resource-saving production and low weight, EPS has a very good environmental balance and is easy to recycle.",
                            "url": "https:\/\/www.novapor.com\/en\/materials\/eps-expanded-polystyrene\/",
                            "breadcrumb": "https:\/\/www.novapor.com › materials › eps-expanded-po..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 25,
                            "rank_absolute": 28,
                            "domain": "www.youtube.com",
                            "title": "Discovery Channel's How It's Made - Expanded Polystyrene ...",
                            "description": "Learn how expanded polystyrene ( EPS ) products are made by the industry's leading manufacturer, ACH Foam Technologies.",
                            "url": "https:\/\/www.youtube.com\/watch?v=XXMque-pLhA",
                            "breadcrumb": "337,2K+ views  ·  13 years ago"
                        },
                        {
                            "type": "organic",
                            "rank_group": 26,
                            "rank_absolute": 29,
                            "domain": "www.tradeinsulations.co.uk",
                            "title": "EPS Insulation - Polystyrene Insulation",
                            "description": "EPS (Expanded Polystyrene) insulation is a kind of foam insulation that is used in both residential and commercial buildings. EPS insulation is very rigid, ...",
                            "url": "https:\/\/www.tradeinsulations.co.uk\/insulation\/material\/eps-insulation\/",
                            "breadcrumb": "https:\/\/www.tradeinsulations.co.uk › insulation › material"
                        },
                        {
                            "type": "organic",
                            "rank_group": 27,
                            "rank_absolute": 30,
                            "domain": "www.thefoamcompany.com.au",
                            "title": "Expanded Polystyrene - EPS - The Foam Company",
                            "description": "Shop high-quality, Australian made Expanded Polystyrene EPS foam. Lightweight, rigid, effective for insulation, packaging, construction, signage, and more.",
                            "url": "https:\/\/www.thefoamcompany.com.au\/collections\/expanded-polystyrene-eps",
                            "breadcrumb": "https:\/\/www.thefoamcompany.com.au › collections › ex..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 28,
                            "rank_absolute": 31,
                            "domain": "www.sciencedirect.com",
                            "title": "Expanded Polystyrene (EPS) geofoam: An introduction to ...",
                            "description": "It encompasses polymeric and non-polymeric foams that are used in geotechnical applications. Geofoams perform functions that traditional geosynthetic products ...",
                            "url": "https:\/\/www.sciencedirect.com\/science\/article\/pii\/0266114494900485",
                            "breadcrumb": "https:\/\/www.sciencedirect.com › science › article › pii"
                        },
                        {
                            "type": "organic",
                            "rank_group": 29,
                            "rank_absolute": 32,
                            "domain": "omnexus.specialchem.com",
                            "title": "Key Applications of Expanded Polystyrene (EPS) - Omnexus",
                            "description": "Popular Applications of Expanded polystyrene (EPS). The features of EPS allow it to be used for a large number of applications - from building to food packaging ...",
                            "url": "https:\/\/omnexus.specialchem.com\/selection-guide\/expanded-polystyrene-eps-foam-insulation\/key-applications",
                            "breadcrumb": "https:\/\/omnexus.specialchem.com › selection-guide › key..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 30,
                            "rank_absolute": 33,
                            "domain": "www.matweb.com",
                            "title": "Overview of materials for Expanded Polystyrene (EPS)",
                            "description": "Material Notes: This property data is a summary of similar materials in the MatWeb database for the category \"Expanded Polystyrene (EPS)\". Each property range ...",
                            "url": "https:\/\/www.matweb.com\/search\/datasheet.aspx?matguid=5f099f2b5eeb41cba804ca0bc64fa62f",
                            "breadcrumb": "https:\/\/www.matweb.com › search › datasheet"
                        },
                        {
                            "type": "organic",
                            "rank_group": 31,
                            "rank_absolute": 34,
                            "domain": "www.baubay.de",
                            "title": "Extruded polystyrene foam, XPS 30 (λ ≥ 0,03, CS ≥ 300)",
                            "description": "The XPS extruded polystyrene is characterised by very good thermal insulation, resistance to humidity impact and high strength. XPS is a homogeneous ...",
                            "url": "https:\/\/www.baubay.de\/extruded-polystyrene-foam-xps-30.html?sl=en",
                            "breadcrumb": "https:\/\/www.baubay.de › ... › Insulation boards"
                        },
                        {
                            "type": "organic",
                            "rank_group": 32,
                            "rank_absolute": 35,
                            "domain": "www.foamex.com.au",
                            "title": "StryroBoard EPS - Expanded Polystyrene",
                            "description": "What is Expanded Polystyrene (EPS)?. Expanded polystyrene is a rigid cellular plastic foam material derived from petroleum and natural gas by-products. It is ...",
                            "url": "https:\/\/www.foamex.com.au\/products\/styroboard-eps",
                            "breadcrumb": "https:\/\/www.foamex.com.au › ... › Credit Application"
                        },
                        {
                            "type": "organic",
                            "rank_group": 33,
                            "rank_absolute": 36,
                            "domain": "www.engineeredfoamproducts.com",
                            "title": "Construction - EPS Foam - Engineered Foam Products",
                            "description": "Expanded Polystyrene (EPS) is a versatile, lightweight alternative to traditional fill materials such as soil or concrete. Due to its high compressive strength ...",
                            "url": "https:\/\/www.engineeredfoamproducts.com\/industries\/construction\/",
                            "breadcrumb": "https:\/\/www.engineeredfoamproducts.com › Industries"
                        },
                        {
                            "type": "organic",
                            "rank_group": 34,
                            "rank_absolute": 37,
                            "domain": "insulationwholesale.co.uk",
                            "title": "EPS Insulation",
                            "description": "EPS (Expanded PolyStyrene) is a white foam plastic substance made from solid polystyrene beads. EPS is a lightweight material with good thermal conductivity ...",
                            "url": "https:\/\/insulationwholesale.co.uk\/insulation\/insulation-board\/polystyrene-insulation-board\/eps-insulation\/",
                            "breadcrumb": "https:\/\/insulationwholesale.co.uk › insulation-board › eps..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 35,
                            "rank_absolute": 38,
                            "domain": "www.isover.co.za",
                            "title": "Expanded Polystyrene",
                            "description": "Expanded Polystyrene (EPS) is a lightweight, rigid, plastic foam insulation material produced from solid beads of polystyrene. EPS contains no CFC's and is ...",
                            "url": "https:\/\/www.isover.co.za\/products\/expanded-polystyrene",
                            "breadcrumb": "https:\/\/www.isover.co.za › products › expanded-polysty..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 36,
                            "rank_absolute": 39,
                            "domain": "www.epstec.com",
                            "title": "What is the expanded polystyrene EPS block process in ...",
                            "description": "What is the expanded polystyrene EPS block process in detail? EPS manufacturing process is highly automated and produces a versatile and ...",
                            "url": "https:\/\/www.epstec.com\/what-is-the-expanded-polystyrene-eps-block-process-in-detail\/",
                            "breadcrumb": "https:\/\/www.epstec.com › what-is-the-expanded-polystyr..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 37,
                            "rank_absolute": 40,
                            "domain": "ultraboard.com",
                            "title": "Extruded Polystyrene vs. Expanded Polystyrene (XPS vs. ...",
                            "description": "Extruded Polystyrene (XPS), is manufactured via an extrusion process where plastic resin and other additives are combined and extruded through a ...",
                            "url": "https:\/\/ultraboard.com\/extruded-polystyrene-vs-expanded-polystyrene-xps-vs-eps-whats-the-difference\/",
                            "breadcrumb": "https:\/\/ultraboard.com › Blog"
                        },
                        {
                            "type": "organic",
                            "rank_group": 38,
                            "rank_absolute": 41,
                            "domain": "supremepetrochem.com",
                            "title": "Expandable polystyrene (EPS) | Producer of EPS in India",
                            "description": "Largest producer of EPS in India · SE250FR. Suitable for medium density block moulding, sandwich panels, insulation boards, 3D panels, cornices and patterns.",
                            "url": "https:\/\/supremepetrochem.com\/eps\/",
                            "breadcrumb": "https:\/\/supremepetrochem.com › eps"
                        },
                        {
                            "type": "organic",
                            "rank_group": 39,
                            "rank_absolute": 42,
                            "domain": "www.linkedin.com",
                            "title": "Risks of Expanded Polystyrene (EPS) Panels",
                            "description": "Toxic Fumes: When EPS panels are exposed to fire or extreme heat, they emit toxic gases, including carbon monoxide and styrene, which can be ...",
                            "url": "https:\/\/www.linkedin.com\/pulse\/risks-expanded-polystyrene-eps-panels-cronossolutions",
                            "breadcrumb": "5 reactions  ·  10 months ago"
                        },
                        {
                            "type": "organic",
                            "rank_group": 40,
                            "rank_absolute": 43,
                            "domain": "newsroom.kunststoffverpackungen.de",
                            "title": "EPS\/polystyrene boxes for fresh fish and ice cream not ...",
                            "description": "Other EPS packaging is expressly not affected by the ban. In a revised draft for guidelines for the uniform interpretation of the directive, the ...",
                            "url": "https:\/\/newsroom.kunststoffverpackungen.de\/en\/2020\/11\/17\/eps-polystyrene-boxes-not-affected-by-ban\/",
                            "breadcrumb": "https:\/\/newsroom.kunststoffverpackungen.de › eps-poly..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 41,
                            "rank_absolute": 44,
                            "domain": "www.insulfoam.com",
                            "title": "What is Expanded Polystyrene?",
                            "description": "Once polystyrene is expanded, a new world of uses expands with it. As EPS foam, the end-product can be used as insulation, protective packaging ...",
                            "url": "https:\/\/www.insulfoam.com\/what-is-expanded-polystyrene-understanding-the-basics-of-eps-foam\/",
                            "breadcrumb": "https:\/\/www.insulfoam.com › what-is-expanded-polysty..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 42,
                            "rank_absolute": 45,
                            "domain": "www.foamex.com.au",
                            "title": "How is Expanded Polystyrene manufactured?",
                            "description": "To create the expanded polystyrene blocks and sheets that you are familiar with, steam is applied to tiny grains or beads of styrene that contain a minute ...",
                            "url": "https:\/\/www.foamex.com.au\/news\/how-is-expanded-polystyrene-manufactured",
                            "breadcrumb": "https:\/\/www.foamex.com.au › ... › Credit Application"
                        },
                        {
                            "type": "organic",
                            "rank_group": 43,
                            "rank_absolute": 46,
                            "domain": "www.foambymail.com",
                            "title": "Polystyrene Foam | Foam Factory, Inc.",
                            "description": "Polystyrene Foam. Expanded Polystyrene (EPS) is a lightweight closed-cell foam used for a wide range of applications, including crafts, insulation, construction ...",
                            "url": "https:\/\/www.foambymail.com\/polystyrene-foam.html",
                            "breadcrumb": "https:\/\/www.foambymail.com › polystyrene-foam"
                        },
                        {
                            "type": "organic",
                            "rank_group": 44,
                            "rank_absolute": 47,
                            "domain": "www.sciencedirect.com",
                            "title": "Expanded polystyrene (EPS) in road construction",
                            "description": "Abstract. Expanded polystyrene (EPS) is a thermoplastic material, derived from pre-expanded polystyrene beads, that combines an extreme lightweight with ...",
                            "url": "https:\/\/www.sciencedirect.com\/science\/article\/pii\/S2352146520301952",
                            "breadcrumb": "https:\/\/www.sciencedirect.com › science › article › pii"
                        },
                        {
                            "type": "organic",
                            "rank_group": 45,
                            "rank_absolute": 48,
                            "domain": "myrealcn.en.made-in-china.com",
                            "title": "2024 EPS Expanded Polystyrene Sheet",
                            "description": "China EPS Expanded Polystyrene Sheet , Insulation Expanded Polystyrene Sandwich Roof Panels EPS Panel Sandwich,1\/5EPS Insulation Board Thermal Insulation ...",
                            "url": "https:\/\/myrealcn.en.made-in-china.com\/product-group\/sqkTIrAJVzVn\/EPS-Expanded-Polystyrene-Sheet-catalog-1.html",
                            "breadcrumb": "https:\/\/myrealcn.en.made-in-china.com › sqkTIrAJVzVn"
                        },
                        {
                            "type": "organic",
                            "rank_group": 46,
                            "rank_absolute": 49,
                            "domain": "www.ambientika.eu",
                            "title": "Built-in casing in expanded polystyrene (EPS) | 100mm",
                            "description": "Description · for precise installation right from the shell construction phase · integrated outward incline · can be shortened · material: expanded polystyrene ...",
                            "url": "https:\/\/www.ambientika.eu\/en\/built-in-casing-in-expanded-polystyrene-eps\/sw10016.1",
                            "breadcrumb": "https:\/\/www.ambientika.eu › built-in-casing-in-expande..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 47,
                            "rank_absolute": 50,
                            "domain": "www.matweb.com",
                            "title": "Overview of materials for Expanded Polystyrene (EPS)",
                            "description": "Tensile Strength, Yield, Average value: 48.4 MPa Grade Count:3 ... Izod Impact, Notched, Average value: 0.235 J\/cm Grade Count:3 ... Insulation R Value, Average ...",
                            "url": "https:\/\/www.matweb.com\/search\/datasheettext.aspx?matguid=5f099f2b5eeb41cba804ca0bc64fa62f",
                            "breadcrumb": "https:\/\/www.matweb.com › search › datasheettext"
                        },
                        {
                            "type": "organic",
                            "rank_group": 48,
                            "rank_absolute": 51,
                            "domain": "www.qualityfoam.com",
                            "title": "EPS Foam (Expanded Polystyrene)",
                            "description": "EPS Foam (Expanded Polystyrene) ... This is an economical packaging foam that is available in densities ranging from 1# to 3# and is easily formed by cutting, hot ...",
                            "url": "http:\/\/www.qualityfoam.com\/expanded-polystyrene.asp",
                            "breadcrumb": "http:\/\/www.qualityfoam.com › expanded-polystyrene"
                        },
                        {
                            "type": "organic",
                            "rank_group": 49,
                            "rank_absolute": 52,
                            "domain": "www.almacenesadda.net",
                            "title": "Properties and Characteristics of Expanded Polystyrene (EPS)",
                            "description": "Expanded polystyrene properties and characteristics · Density · Color · Mechanical strength · Thermal insulation · Water and water vapour behaviour.",
                            "url": "http:\/\/www.almacenesadda.net\/en\/blog\/properties-and-characteristics-of-expanded-polystyrene-eps\/",
                            "breadcrumb": "http:\/\/www.almacenesadda.net › blog › properties-and-c..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 50,
                            "rank_absolute": 53,
                            "domain": "www.linkedin.com",
                            "title": "Manufacturing Processes of Expandable Polystyrene (EPS)",
                            "description": "Expandable polystyrene (EPS) is produced in a suspension process by adding a blowing agent, usually pentane, which causes resin to foam during ...",
                            "url": "https:\/\/www.linkedin.com\/pulse\/manufacturing-processes-expandable-polystyrene-eps-david-zhang",
                            "breadcrumb": "20+ reactions  ·  3 years ago"
                        },
                        {
                            "type": "organic",
                            "rank_group": 51,
                            "rank_absolute": 54,
                            "domain": "www.ita.rwth-aachen.de",
                            "title": "Innovative use of expanded polystyrene (EPS) for textile ...",
                            "description": "Innovative use of expanded polystyrene (EPS) for textile preforming · loading.",
                            "url": "https:\/\/www.ita.rwth-aachen.de\/custom\/rwth\/desktop.asp?url=%2Fcms%2FITA%2FForschung%2FPublikationen%2F~jfbg%2FDetails%2Flidx%2F1%2F%3Ffile%3D225956",
                            "breadcrumb": "https:\/\/www.ita.rwth-aachen.de › custom › rwth › desktop"
                        },
                        {
                            "type": "organic",
                            "rank_group": 52,
                            "rank_absolute": 55,
                            "domain": "www.designingbuildings.co.uk",
                            "title": "Expanded polystyrene",
                            "description": "Often used in applications where moisture resistance and some strength is important such as below ground or below slabs. It is made by heating ...",
                            "url": "https:\/\/www.designingbuildings.co.uk\/wiki\/Expanded%20polystyrene",
                            "breadcrumb": "https:\/\/www.designingbuildings.co.uk › wiki › Expanded..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 53,
                            "rank_absolute": 56,
                            "domain": "www.usimm.ca",
                            "title": "Should You Choose Expanded or Extruded Polystyrene",
                            "description": "The Difference between Expanded Polystyrene and Extruded Polystyrene. The main difference between these two materials lies in their density.",
                            "url": "https:\/\/www.usimm.ca\/en\/polystyrene-panels-should-you-choose-expanded-or-extruded-polystyrene\/",
                            "breadcrumb": "https:\/\/www.usimm.ca › polystyrene-panels-should-you-..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 54,
                            "rank_absolute": 57,
                            "domain": "www.eps.co.uk",
                            "title": "The BPF Expanded Polystyrene Group | What is EPS made ...",
                            "description": "The BPF EPS (expanded polystyrene) Group is a member-group, representing 80% of the EPS manufacturing industry in the UK. It works to provide authoritative, ...",
                            "url": "https:\/\/www.eps.co.uk\/abouteps\/composition.html",
                            "breadcrumb": "https:\/\/www.eps.co.uk › abouteps › composition"
                        },
                        {
                            "type": "organic",
                            "rank_group": 55,
                            "rank_absolute": 58,
                            "domain": "geofoamintl.com",
                            "title": "What Is Expanded Polystyrene (EPS)?",
                            "description": "EPS geofoam is 98% to 99% air by volume. The air pockets that are formed during the manufacturing process make it a great thermal insulator.",
                            "url": "https:\/\/geofoamintl.com\/what-is-expanded-polystyrene-eps\/",
                            "breadcrumb": "https:\/\/geofoamintl.com › Blog"
                        },
                        {
                            "type": "organic",
                            "rank_group": 56,
                            "rank_absolute": 59,
                            "domain": "www.youtube.com",
                            "title": "What is HD (High Density) EPS Foam? Expanded Polystyrene ...",
                            "description": "http:\/\/www.EPSFoamPro.com Basically HD foam stands for High Density Foam. EPS Foam in Manufactured in Many Different Densities, and some EPS ...",
                            "url": "https:\/\/www.youtube.com\/watch?v=Au4M-R33ZLY",
                            "breadcrumb": "26K+ views  ·  14 years ago"
                        },
                        {
                            "type": "organic",
                            "rank_group": 57,
                            "rank_absolute": 60,
                            "domain": "www.inokor.co",
                            "title": "Expanded Polystyrene technology - EPS",
                            "description": "EPS is a lightweight insulation material with many attractive properties. It is produced by applying heat and water to granulated Polystyrene. The expanded ...",
                            "url": "https:\/\/www.inokor.co\/eu\/technology-eps\/",
                            "breadcrumb": "https:\/\/www.inokor.co › technology-eps"
                        },
                        {
                            "type": "organic",
                            "rank_group": 58,
                            "rank_absolute": 61,
                            "domain": "epsfoam.co.nz",
                            "title": "EPS Foam NZ: Expanded Polystyrene Foam | Foam Insulation",
                            "description": "EPS Form provides insulation, insulated panels, void foam, spacer, reinforcing products, pods, CNC, polystyrene sheets, rolls and blocks.",
                            "url": "https:\/\/epsfoam.co.nz\/",
                            "breadcrumb": "https:\/\/epsfoam.co.nz"
                        },
                        {
                            "type": "organic",
                            "rank_group": 59,
                            "rank_absolute": 62,
                            "domain": "smartpackagingeurope.eu",
                            "title": "EPS at a glance - Smart Packaging Europe",
                            "description": "EPS packaging is so light that some call it 'engineered air'. Invented in 1949 by Dr. Fritz Stastny, a scientist working at BASF in Germany, EPS has been used ...",
                            "url": "https:\/\/smartpackagingeurope.eu\/eps-at-a-glance\/",
                            "breadcrumb": "https:\/\/smartpackagingeurope.eu › eps-at-a-glance"
                        },
                        {
                            "type": "organic",
                            "rank_group": 60,
                            "rank_absolute": 63,
                            "domain": "www.polyform.com",
                            "title": "Expanded Polystyrene and Water - Polyform",
                            "description": "EPS is resistant to water and moisture. This means that it does not absorb them and that its insulating performance is not compromised when in ...",
                            "url": "https:\/\/www.polyform.com\/uncategorized\/le-polystyrene-expanse-et-leau\/?lang=en",
                            "breadcrumb": "https:\/\/www.polyform.com › uncategorized › le-polysty..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 61,
                            "rank_absolute": 64,
                            "domain": "greatwhitefoam.ca",
                            "title": "Expanded Polystyrene (EPS) Advantages",
                            "description": "Economical. Expanded Polystyrene (EPS) is more economical than extruded polystyrene (XPS) and polyisocyanurate (PolyIso) products. EPS products cost less when ...",
                            "url": "https:\/\/greatwhitefoam.ca\/eps-advantages",
                            "breadcrumb": "https:\/\/greatwhitefoam.ca › eps-advantages"
                        },
                        {
                            "type": "organic",
                            "rank_group": 62,
                            "rank_absolute": 65,
                            "domain": "www.hsv-pi.nl",
                            "title": "EPS Expanded Polystyrene - HSV",
                            "description": "Fritz Stastny, a scientist working at BASF in Germany, invented expanding polystyrene, in short: EPS. When the pearl-shaped raw material is exposed to steam it ...",
                            "url": "https:\/\/www.hsv-pi.nl\/en\/eps\/",
                            "breadcrumb": "https:\/\/www.hsv-pi.nl › eps"
                        },
                        {
                            "type": "organic",
                            "rank_group": 63,
                            "rank_absolute": 66,
                            "domain": "michiganfoam.com",
                            "title": "Expanded Polystyrene | Michigan Foam Products LLC",
                            "description": "Introduction. Expanded Polystyrene, or EPS, goes through a series of changes that take it from a small polystyrene pre-expanded bead about 1mm in width to an ...",
                            "url": "https:\/\/michiganfoam.com\/expanded-polystyrene\/",
                            "breadcrumb": "https:\/\/michiganfoam.com › expanded-polystyrene"
                        },
                        {
                            "type": "organic",
                            "rank_group": 64,
                            "rank_absolute": 67,
                            "domain": "polymers.totalenergies.com",
                            "title": "Insulation - Polymers | - TotalEnergies.com",
                            "description": "EPS (Expandable PolyStyrene) or airpop® has been used for over 70 years for many purposes. It was originally intended as insulation material and it is still ...",
                            "url": "https:\/\/polymers.totalenergies.com\/markets\/infrastructure-construction\/insulation",
                            "breadcrumb": "https:\/\/polymers.totalenergies.com › markets › insulation"
                        },
                        {
                            "type": "organic",
                            "rank_group": 65,
                            "rank_absolute": 68,
                            "domain": "www.insulfoam.com",
                            "title": "EPS Insulation Products",
                            "description": "Insulfoam HD Composite (InsulFoam HD Composite) is an engineered insulation consisting of a closed-cell, lightweight and resilient expanded polystyrene (EPS) ...",
                            "url": "https:\/\/www.insulfoam.com\/insulfoam-eps\/",
                            "breadcrumb": "https:\/\/www.insulfoam.com › insulfoam-eps"
                        },
                        {
                            "type": "organic",
                            "rank_group": 66,
                            "rank_absolute": 69,
                            "domain": "epsa.org.au",
                            "title": "Expanded Polystyrene Australia: Home",
                            "description": "Lightweight Wonder. Expanded Polystyrene (EPS) is a lightweight cellular plastic material consisting of small hollow spherical balls. It is this closed cellular ...",
                            "url": "https:\/\/epsa.org.au\/",
                            "breadcrumb": "https:\/\/epsa.org.au"
                        },
                        {
                            "type": "organic",
                            "rank_group": 67,
                            "rank_absolute": 70,
                            "domain": "www.thomasnet.com",
                            "title": "Expanded Polystyrene (EPS): Properties and Applications",
                            "description": "With a range of compressive strengths, EPS functions as a stiff, lightweight insulation material that can withstand loads and backfill forces.",
                            "url": "https:\/\/www.thomasnet.com\/articles\/plastics-rubber\/expanded-polystyrene-eps\/",
                            "breadcrumb": "https:\/\/www.thomasnet.com › articles › plastics-rubber"
                        },
                        {
                            "type": "organic",
                            "rank_group": 68,
                            "rank_absolute": 71,
                            "domain": "www.nuclear-power.com",
                            "title": "Expanded Polystyrene - EPS - Thermal Insulation",
                            "description": "The lower the thermal conductivity of the material, the greater the material's ability to resist heat transfer, and hence the greater the insulation's ...",
                            "url": "https:\/\/www.nuclear-power.com\/nuclear-engineering\/heat-transfer\/heat-losses\/insulation-materials\/expanded-polystyrene-eps\/",
                            "breadcrumb": "https:\/\/www.nuclear-power.com › insulation-materials"
                        },
                        {
                            "type": "organic",
                            "rank_group": 69,
                            "rank_absolute": 72,
                            "domain": "medium.com",
                            "title": "Why Styrofoam (Expanded Polystyrene) Should Be ...",
                            "description": "“Expanded Polystyrene waste poses a risk to the fragile ecological balance, since marine and land wildlife often perish as a result of ingesting ...",
                            "url": "https:\/\/medium.com\/age-of-awareness\/why-styrofoam-expanded-polystyrene-should-be-banned-everywhere-in-the-world-4101552f5e2b",
                            "breadcrumb": "140+ likes  ·  5 years ago"
                        },
                        {
                            "type": "organic",
                            "rank_group": 70,
                            "rank_absolute": 73,
                            "domain": "paneltech.eu",
                            "title": "Expanded polystyrene – paneltech.eu",
                            "description": "Expanded polystyrene Paneltech is regarded as an excellent insulation material, applied commonly in building industry.",
                            "url": "https:\/\/paneltech.eu\/eps",
                            "breadcrumb": "https:\/\/paneltech.eu › eps"
                        },
                        {
                            "type": "organic",
                            "rank_group": 71,
                            "rank_absolute": 74,
                            "domain": "de.pons.com",
                            "title": "expanded polystyrene - Englisch-Deutsch Übersetzung",
                            "description": "Einsprachige Beispiele (nicht von der PONS Redaktion geprüft). Englisch. In recent years the expanded polystyrene composites with cellulose and starch have also ...",
                            "url": "https:\/\/de.pons.com\/%C3%BCbersetzung\/englisch-deutsch\/expanded+polystyrene",
                            "breadcrumb": "https:\/\/de.pons.com › übersetzung"
                        },
                        {
                            "type": "organic",
                            "rank_group": 72,
                            "rank_absolute": 75,
                            "domain": "markofoam.com",
                            "title": "Expanded Polystyrene (EPS)",
                            "description": "EPS (Expanded Polystyrene) is a closed-cell rigid foam and one of the most cost-effective materials on the market. Its wide range of properties make it an ...",
                            "url": "https:\/\/markofoam.com\/pages\/expanded-polystyrene-eps",
                            "breadcrumb": "https:\/\/markofoam.com › pages › expanded-polystyrene-..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 73,
                            "rank_absolute": 76,
                            "domain": "www.eps.co.uk",
                            "title": "The BPF Expanded Polystyrene (EPS) Group",
                            "description": "The BPF EPS (expanded polystyrene) Group is a member-group, representing 80% of the EPS manufacturing industry in the UK. It works to provide authoritative, ...",
                            "url": "https:\/\/www.eps.co.uk\/",
                            "breadcrumb": "https:\/\/www.eps.co.uk"
                        },
                        {
                            "type": "organic",
                            "rank_group": 74,
                            "rank_absolute": 77,
                            "domain": "www.bpf.co.uk",
                            "title": "Moulding Expanded Polystyrene (EPS)",
                            "description": "The tiny spherical polystyrene beads are expanded to about 40 times their original size using a small quantity of pentane (typically 5% by weight) as a blowing ...",
                            "url": "https:\/\/www.bpf.co.uk\/plastipedia\/processes\/Moulding_EPS.aspx",
                            "breadcrumb": "https:\/\/www.bpf.co.uk › plastipedia › processes › Moul..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 75,
                            "rank_absolute": 78,
                            "domain": "hotwiresystems.com",
                            "title": "Polystyrene as a versatile material! | Different uses of ...",
                            "description": "The foam plastic consists of 96 – 98% air and 2-4% polystyrene. Granules are heated with steam in such a way that they are rapidly expanding (foaming) and ...",
                            "url": "https:\/\/hotwiresystems.com\/what-is-polystyrene-eps-xps-different-uses-of-polystyrene\/",
                            "breadcrumb": "https:\/\/hotwiresystems.com › what-is-polystyrene-eps-xps..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 76,
                            "rank_absolute": 79,
                            "domain": "www.insulationshop.co",
                            "title": "Polystyrene Boards",
                            "description": "Polystyrene is a rigid foam insulation, which has an exceptional ability to insulate against noise and extreme temperatures and also waterproof and ...",
                            "url": "https:\/\/www.insulationshop.co\/polystyrene_insulation_eps_70_online.html",
                            "breadcrumb": "https:\/\/www.insulationshop.co › polystyrene_insulation_..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 77,
                            "rank_absolute": 80,
                            "domain": "warm-international.com",
                            "title": "Expanded Polystyrene (EPS) Insulation Panels",
                            "description": "These panels, which are used for insulation, provide heat and sound insulation. It preserves cold air or hot air in accordance with the season in every season ...",
                            "url": "https:\/\/warm-international.com\/products\/eps\/",
                            "breadcrumb": "https:\/\/warm-international.com › Products"
                        },
                        {
                            "type": "organic",
                            "rank_group": 78,
                            "rank_absolute": 81,
                            "domain": "www.linguee.com",
                            "title": "expanded polystyrene - German translation",
                            "description": "expanded polystyrene noun— · Styropor nt · See also: expanded —.",
                            "url": "https:\/\/www.linguee.com\/english-german\/translation\/expanded+polystyrene.html",
                            "breadcrumb": "https:\/\/www.linguee.com › english-german › expanded..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 79,
                            "rank_absolute": 82,
                            "domain": "www.eymenpetrokimya.com.tr",
                            "title": "EPS - Expanded Polystyrene | Eymen Petrochemicals",
                            "description": "When the polystyrene raw material contacts the water vapor, the pentane gas granules within the material are expanded. When these expanded gas granules adhere ...",
                            "url": "https:\/\/www.eymenpetrokimya.com.tr\/en\/products\/eps-expanded-polystyrene\/",
                            "breadcrumb": "https:\/\/www.eymenpetrokimya.com.tr › ... › Products"
                        },
                        {
                            "type": "organic",
                            "rank_group": 80,
                            "rank_absolute": 83,
                            "domain": "www.youtube.com",
                            "title": "EPS (Expanded Polystyrene) Explainer Video",
                            "description": "Why EPS ? Lightweight, shock-resistant, and an exceptional insulator: EPS is a highly versatile packaging material and the film shows you how ...",
                            "url": "https:\/\/www.youtube.com\/watch?v=yfl3AbO-tAU",
                            "breadcrumb": "3K+ views  ·  2 years ago"
                        },
                        {
                            "type": "organic",
                            "rank_group": 81,
                            "rank_absolute": 84,
                            "domain": "www.koolfoam.com.au",
                            "title": "Expanded Polystyrene (EPS) Sheets and Insulation - Koolfoam",
                            "description": "The Koolfoam Ultra polystyrene sheets are perfect for insulations with a rating of R 3.57 for 100mm compared to regular EPS with a rating of R 2.78.",
                            "url": "https:\/\/www.koolfoam.com.au\/eps-sheets\/",
                            "breadcrumb": "https:\/\/www.koolfoam.com.au › eps-sheets"
                        },
                        {
                            "type": "organic",
                            "rank_group": 82,
                            "rank_absolute": 85,
                            "domain": "www.epsindustry.org",
                            "title": "EPS Industry Alliance",
                            "description": "The EPS Industry Alliance (EPS-IA) is the North American trade association for the expanded polystyrene (EPS) industry. Our members – more than 50 ...",
                            "url": "https:\/\/www.epsindustry.org\/",
                            "breadcrumb": "https:\/\/www.epsindustry.org"
                        },
                        {
                            "type": "organic",
                            "rank_group": 83,
                            "rank_absolute": 86,
                            "domain": "www.buyinsulationonline.co.uk",
                            "title": "Blogs Facts About Expanded Polystyrene Insulation",
                            "description": "Expanded polystyrene boards are completely safe to use in all construction applications. EPS foam boards are non-toxic, odourless, chemically inert and non- ...",
                            "url": "https:\/\/www.buyinsulationonline.co.uk\/blog\/facts-about-expanded-polystyrene-insulation",
                            "breadcrumb": "https:\/\/www.buyinsulationonline.co.uk › blog › facts-abo..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 84,
                            "rank_absolute": 87,
                            "domain": "carpenter.com",
                            "title": "Expanded Polystyrene Systems",
                            "description": "Carpenter EPS Roof Insulation provides an excellent thermal barrier for both new construction and reproofing applications. Our UL certified EPS ...",
                            "url": "https:\/\/carpenter.com\/en\/divisions\/expanded-polystyrene-systems\/",
                            "breadcrumb": "https:\/\/carpenter.com › divisions › expanded-polystyren..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 85,
                            "rank_absolute": 88,
                            "domain": "www.madehow.com",
                            "title": "Expanded Polystyrene Foam (EPF)",
                            "description": "Polystyrene is made in a process known as suspension polymerization. After styrene ...",
                            "url": "https:\/\/www.madehow.com\/Volume-1\/Expanded-Polystyrene-Foam-EPF.html",
                            "breadcrumb": "https:\/\/www.madehow.com › Volume-1 › Expanded-Pol..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 86,
                            "rank_absolute": 89,
                            "domain": "www.baunetzwissen.de",
                            "title": "Expandiertes Polystyrol (EPS) | Dämmstoffe",
                            "description": "Expandiertes Polystyrol (EPS). Expandiertes Polystyrol gehört zur Gruppe der organischen, synthetischen Dämmstoffe. EPS-Dämmstoffe können für nahezu jede ...",
                            "url": "https:\/\/www.baunetzwissen.de\/daemmstoffe\/fachwissen\/daemmstoffe\/expandiertes-polystyrol-eps-152198",
                            "breadcrumb": "https:\/\/www.baunetzwissen.de › ex..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 87,
                            "rank_absolute": 91,
                            "domain": "www.grandviewresearch.com",
                            "title": "Expanded Polystyrene Market Share Analysis Report, 2030",
                            "description": "The global expanded polystyrene market was valued at USD 10.18 billion in 2022 and is anticipated to expand at a compound annual growth rate (CAGR) of 8.59% ...",
                            "url": "https:\/\/www.grandviewresearch.com\/industry-analysis\/expanded-polystyrene-eps-market",
                            "breadcrumb": "https:\/\/www.grandviewresearch.com › industry-analysis"
                        },
                        {
                            "type": "organic",
                            "rank_group": 88,
                            "rank_absolute": 92,
                            "domain": "www.marketsandmarkets.com",
                            "title": "The Global Expanded Polystyrene Market, Industry Size ...",
                            "description": "APAC has emerged as the leading consumer and producer of EPS due to the increasing demand from the domestic front and rising income levels. The easy ...",
                            "url": "https:\/\/www.marketsandmarkets.com\/Market-Reports\/expanded-polystyrene-market-1138.html",
                            "breadcrumb": "https:\/\/www.marketsandmarkets.com › Market-Reports"
                        },
                        {
                            "type": "organic",
                            "rank_group": 89,
                            "rank_absolute": 93,
                            "domain": "onlinelibrary.wiley.com",
                            "title": "Application of expanded polystyrene (EPS) in buildings ...",
                            "description": "EPS foam is used as core to maintain the vacuum condition as well as to provide support for the envelope. The desiccant is placed in the core as ...",
                            "url": "https:\/\/onlinelibrary.wiley.com\/doi\/10.1002\/app.47529",
                            "breadcrumb": "https:\/\/onlinelibrary.wiley.com › doi › app"
                        },
                        {
                            "type": "organic",
                            "rank_group": 90,
                            "rank_absolute": 94,
                            "domain": "www.sulzer.com",
                            "title": "Expandierbare Polystyrole (EPS)",
                            "description": "Expandierbare Polystyrole (ESP) bestehen aus Polystyrol-Mikropellets oder -kugeln, die Treibmittel oder andere Additive zum Aufschäumen enthalten.",
                            "url": "https:\/\/www.sulzer.com\/de-ch\/germany\/shared\/products\/expandable-polystyrene-eps",
                            "breadcrumb": "https:\/\/www.sulzer.com › products"
                        },
                        {
                            "type": "organic",
                            "rank_group": 91,
                            "rank_absolute": 95,
                            "domain": "insulationcorp.com",
                            "title": "The Difference Between STYROFOAM™ and EPS",
                            "description": "Those large white foam blocks are not \"Styrofoam Blocks\", they're EPS, which is more versatile & cost effective than XPS!",
                            "url": "https:\/\/insulationcorp.com\/the-difference-between-styrofoam-and-eps\/",
                            "breadcrumb": "https:\/\/insulationcorp.com › the-difference-between-styr..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 92,
                            "rank_absolute": 96,
                            "domain": "insultherme.com",
                            "title": "Expanded Polystyrene (EPS) Insulation | Pre insulated Duct",
                            "description": "Insultherm® Expanded Polystyrene (EPS) insulation offers cost-effectiveness and energy efficiency. It's thermal and mechanical properties are ideal for ...",
                            "url": "https:\/\/insultherme.com\/expanded-polystyrene-insulation.html",
                            "breadcrumb": "https:\/\/insultherme.com › expanded-polystyrene-insulati..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 93,
                            "rank_absolute": 97,
                            "domain": "eumeps.eu",
                            "title": "Home - Eumeps Corporate - The voice of the European EPS ...",
                            "description": "EPS in Numbers · A dedicated effort to collect EPS construction cut-offs · For decades, EPS has been facilitating people's lives · EPS embraces the future of ...",
                            "url": "https:\/\/eumeps.eu\/",
                            "breadcrumb": "https:\/\/eumeps.eu"
                        },
                        {
                            "type": "organic",
                            "rank_group": 94,
                            "rank_absolute": 98,
                            "domain": "www.dcceew.gov.au",
                            "title": "National Plastics Plan - Pathway to more sustainable use ...",
                            "description": "National Plastics Plan - Pathway to more sustainable use of expanded polystyrene (EPS) · EPS phase outs · Why are certain EPS packaging products ...",
                            "url": "https:\/\/www.dcceew.gov.au\/environment\/protection\/waste\/publications\/npp-pathway-to-more-sustainable-use-eps",
                            "breadcrumb": "https:\/\/www.dcceew.gov.au › waste › publications › np..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 95,
                            "rank_absolute": 99,
                            "domain": "www.spglobal.com",
                            "title": "Expandable Polystyrene - Chemical Economics Handbook ...",
                            "description": "Expandable polystyrene (EPS) is a rigid cellular plastic originally invented in Germany by BASF in 1950. EPS is 98% air and the rest is made from tiny, ...",
                            "url": "https:\/\/www.spglobal.com\/commodityinsights\/en\/ci\/products\/expandable-polystyrene-chemical-economics-handbook.html",
                            "breadcrumb": "https:\/\/www.spglobal.com › commodityinsights › products"
                        },
                        {
                            "type": "organic",
                            "rank_group": 96,
                            "rank_absolute": 100,
                            "domain": "iwi-gmbh.com",
                            "title": "What is EPS (expanded polystyrene particle foam):",
                            "description": "In the case of expanded polystyrene particle foam (EPS) the polystyrene granulate (polystyrene powder), into which the foaming agent pentane has been ...",
                            "url": "http:\/\/iwi-gmbh.com\/iwi\/pdf\/IWI_EPS_en.pdf",
                            "breadcrumb": "http:\/\/iwi-gmbh.com › iwi › pdf › IWI_EPS_en"
                        },
                        {
                            "type": "organic",
                            "rank_group": 97,
                            "rank_absolute": 101,
                            "domain": "fabcopolystyrene.com",
                            "title": "EXPANDED POLYSTYRENE (EPS)",
                            "description": "EXPANDED POLYSTYRENE (EPS). FABCO BLOCK EPS ( SPECIAL CUT INSULATED VOLCANIC BLOCK ). Quick View. EXPANDED POLYSTYRENE (EPS). Fabco Board Roof ( EPS ).",
                            "url": "https:\/\/fabcopolystyrene.com\/product-category\/expanded-polystyrene-en\/",
                            "breadcrumb": "https:\/\/fabcopolystyrene.com › product-category › expa..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 98,
                            "rank_absolute": 102,
                            "domain": "businessanalytiq.com",
                            "title": "Expandable Polystyrene (EPS) price index",
                            "description": "Expandable Polystyrene (EPS) price July 2024 and outlook (see chart below) · North America:US$1.77\/KG, 2.3% up · Europe:US$2.04\/KG, -7.3% down · Northeast Asia: ...",
                            "url": "https:\/\/businessanalytiq.com\/procurementanalytics\/index\/expandable-polystyrene-eps-price-index\/",
                            "breadcrumb": "https:\/\/businessanalytiq.com › procurementanalytics › e..."
                        },
                        {
                            "type": "organic",
                            "rank_group": 99,
                            "rank_absolute": 103,
                            "domain": "insutech-eg.com",
                            "title": "What are Expanded Polystyrene - EPS Foam Blocks?",
                            "description": "– EPS Foam Blocks provide a raised structural floor over the concrete slab, in a short time frame which speeds up the raised floor construction ...",
                            "url": "https:\/\/insutech-eg.com\/what-are-expanded-polystyrene-eps-foam-blocks\/",
                            "breadcrumb": "https:\/\/insutech-eg.com › what-are-expanded-polystyre..."
                        }
                    ]
                }
            ]
        }
    ]
}
*/
