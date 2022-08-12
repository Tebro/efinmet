package main

const airportsJSON = `{
  "EFAA": [
    {
      "lat": 67.603611,
      "lon": 23.971667,
      "name": "AAVAHELUKKA"
    }
  ],
  "EFAH": [
    {
      "lat": 64.895278,
      "lon": 25.752222,
      "name": "AHMOSUO"
    }
  ],
  "EFAL": [
    {
      "lat": 62.554722,
      "lon": 23.573333,
      "name": "ALAVUS"
    }
  ],
  "EFEJ": [
    {
      "lat": 60.2208333,
      "lon": 24.6863889,
      "name": "JORVIN SAIRAALA"
    }
  ],
  "EFEK": [
    {
      "lat": 69.0022222,
      "lon": 20.8963889,
      "name": "KILPISJÄRVI"
    }
  ],
  "EFET": [
    {
      "lat": 68.364444,
      "lon": 23.4275,
      "name": "ENONTEKIÖ"
    }
  ],
  "EFEU": [
    {
      "lat": 61.116111,
      "lon": 22.201389,
      "name": "EURA"
    }
  ],
  "EFFH": [
    {
      "lat": 62.9244444,
      "lon": 27.7233333,
      "name": "KELLONIEMI"
    }
  ],
  "EFFO": [
    {
      "lat": 60.801944,
      "lon": 23.645556,
      "name": "FORSSA"
    }
  ],
  "EFGE": [
    {
      "lat": 60.086944,
      "lon": 22.521944,
      "name": "GENBÖLE"
    }
  ],
  "EFHA": [
    {
      "lat": 61.856389,
      "lon": 24.789167,
      "name": "HALLI"
    }
  ],
  "EFHF": [
    {
      "lat": 60.253889,
      "lon": 25.044167,
      "name": "HELSINKI-MALMI"
    }
  ],
  "EFHK": [
    {
      "lat": 60.317222,
      "lon": 24.963333,
      "name": "HELSINKI-VANTAA"
    }
  ],
  "EFHL": [
    {
      "lat": 64.969722,
      "lon": 24.704722,
      "name": "HAILUOTO"
    }
  ],
  "EFHM": [
    {
      "lat": 61.690278,
      "lon": 23.074167,
      "name": "HÄMEENKYRÖ"
    }
  ],
  "EFHN": [
    {
      "lat": 59.845556,
      "lon": 23.082222,
      "name": "HANKO"
    }
  ],
  "EFHO": [
    {
      "lat": 65.0013889,
      "lon": 25.5186111,
      "name": "OULUN YLIOPISTOLLINEN SAIRAALA"
    }
  ],
  "EFHP": [
    {
      "lat": 64.113056,
      "lon": 25.504167,
      "name": "HAAPAVESI"
    }
  ],
  "EFHS": [
    {
      "lat": 62.7722222,
      "lon": 22.8188889,
      "name": "SEINÄJOEN KESKUSSAIRAALA"
    }
  ],
  "EFHV": [
    {
      "lat": 60.654444,
      "lon": 24.881111,
      "name": "HYVINKÄÄ"
    }
  ],
  "EFHY": [
    {
      "lat": 60.1888889,
      "lon": 24.9075,
      "name": "HELSINGIN YLIOPISTOLLINEN KESKUSSAIRAALA/MEILAHTI"
    }
  ],
  "EFII": [
    {
      "lat": 63.631944,
      "lon": 27.122222,
      "name": "IISALMI"
    }
  ],
  "EFIK": [
    {
      "lat": 60.462778,
      "lon": 23.651111,
      "name": "KIIKALA"
    }
  ],
  "EFIM": [
    {
      "lat": 61.249722,
      "lon": 28.903611,
      "name": "IMMOLA"
    }
  ],
  "EFIT": [
    {
      "lat": 62.166111,
      "lon": 30.073611,
      "name": "KITEE"
    }
  ],
  "EFIV": [
    {
      "lat": 68.610833,
      "lon": 27.413889,
      "name": "IVALO"
    }
  ],
  "EFJE": [
    {
      "lat": 62.5908333,
      "lon": 29.7777778,
      "name": "POHJOIS-KARJALAN KESKUSSAIRAALA"
    }
  ],
  "EFJM": [
    {
      "lat": 61.778611,
      "lon": 22.716111,
      "name": "JÄMIJÄRVI"
    }
  ],
  "EFJO": [
    {
      "lat": 62.658889,
      "lon": 29.624444,
      "name": "JOENSUU"
    }
  ],
  "EFJV": [
    {
      "lat": 62.2297222,
      "lon": 25.7041667,
      "name": "KESKI-SUOMEN KESKUSSAIRAALA NOVA"
    }
  ],
  "EFJY": [
    {
      "lat": 62.400833,
      "lon": 25.672778,
      "name": "JYVÄSKYLÄ"
    }
  ],
  "EFKA": [
    {
      "lat": 63.124167,
      "lon": 23.051389,
      "name": "KAUHAVA"
    }
  ],
  "EFKE": [
    {
      "lat": 65.779167,
      "lon": 24.584722,
      "name": "KEMI/TORNIO"
    }
  ],
  "EFKG": [
    {
      "lat": 60.246944,
      "lon": 20.804722,
      "name": "KUMLINGE"
    }
  ],
  "EFKH": [
    {
      "lat": 64.1125,
      "lon": 29.438611,
      "name": "KUHMO"
    }
  ],
  "EFKI": [
    {
      "lat": 64.284167,
      "lon": 27.6875,
      "name": "KAJAANI"
    }
  ],
  "EFKJ": [
    {
      "lat": 62.463333,
      "lon": 22.391111,
      "name": "KAUHAJOKI"
    }
  ],
  "EFKK": [
    {
      "lat": 63.720278,
      "lon": 23.139167,
      "name": "KOKKOLA/PIETARSAARI"
    }
  ],
  "EFKM": [
    {
      "lat": 66.715833,
      "lon": 27.157222,
      "name": "KEMIJARVI"
    }
  ],
  "EFKO": [
    {
      "lat": 64.228611,
      "lon": 23.826389,
      "name": "KALAJOKI"
    }
  ],
  "EFKR": [
    {
      "lat": 63.989167,
      "lon": 25.743611,
      "name": "KARSAMAKI"
    }
  ],
  "EFKS": [
    {
      "lat": 65.990278,
      "lon": 29.231944,
      "name": "KUUSAMO"
    }
  ],
  "EFKT": [
    {
      "lat": 67.698611,
      "lon": 24.848056,
      "name": "KITTILA"
    }
  ],
  "EFKU": [
    {
      "lat": 63.008611,
      "lon": 27.794444,
      "name": "KUOPIO"
    }
  ],
  "EFKV": [
    {
      "lat": 63.125278,
      "lon": 25.124167,
      "name": "KIVIJARVI"
    }
  ],
  "EFKY": [
    {
      "lat": 60.571389,
      "lon": 26.896111,
      "name": "KYMI"
    }
  ],
  "EFLA": [
    {
      "lat": 61.144167,
      "lon": 25.693056,
      "name": "LAHTI/VESIVEHMAA"
    }
  ],
  "EFLL": [
    {
      "lat": 63.399444,
      "lon": 27.478889,
      "name": "LAPINLAHTI"
    }
  ],
  "EFLN": [
    {
      "lat": 63.511944,
      "lon": 29.629167,
      "name": "LIEKSA-NURMES"
    }
  ],
  "EFLP": [
    {
      "lat": 61.045833,
      "lon": 28.148611,
      "name": "LAPPEENRANTA"
    }
  ],
  "EFMA": [
    {
      "lat": 60.121944,
      "lon": 19.896389,
      "name": "MARIEHAMN"
    }
  ],
  "EFME": [
    {
      "lat": 62.946667,
      "lon": 23.518889,
      "name": "MENKIJARVI"
    }
  ],
  "EFMH": [
    {
      "lat": 60.1147222,
      "lon": 19.9330556,
      "name": "AHVENANMAAN KESKUSSAIRAALA"
    }
  ],
  "EFMI": [
    {
      "lat": 61.686389,
      "lon": 27.2,
      "name": "MIKKELI"
    }
  ],
  "EFMN": [
    {
      "lat": 60.5725,
      "lon": 25.508889,
      "name": "MANTSALA"
    }
  ],
  "EFMS": [
    {
      "lat": 61.6925,
      "lon": 27.2805556,
      "name": "MIKKELIN SAIRAALAPARKKI"
    }
  ],
  "EFNS": [
    {
      "lat": 60.52,
      "lon": 24.831667,
      "name": "SAVIKKO"
    }
  ],
  "EFNU": [
    {
      "lat": 60.333889,
      "lon": 24.296389,
      "name": "NUMMELA"
    }
  ],
  "EFOP": [
    {
      "lat": 60.876389,
      "lon": 22.744722,
      "name": "ORIPAA"
    }
  ],
  "EFOU": [
    {
      "lat": 64.929167,
      "lon": 25.355556,
      "name": "OULU"
    }
  ],
  "EFPA": [
    {
      "lat": 68.145556,
      "lon": 25.822778,
      "name": "POKKA"
    }
  ],
  "EFPE": [
    {
      "lat": 60.3311111,
      "lon": 25.0608333,
      "name": "PEIJAKSEN SAIRAALA"
    }
  ],
  "EFPI": [
    {
      "lat": 61.245278,
      "lon": 22.195,
      "name": "PIIKAJARVI"
    }
  ],
  "EFPJ": [
    {
      "lat": 62.8975,
      "lon": 27.6483333,
      "name": "KUOPION YLIOPISTOLLINEN SAIRAALA"
    }
  ],
  "EFPK": [
    {
      "lat": 62.264722,
      "lon": 27.002778,
      "name": "PIEKSAMAKI"
    }
  ],
  "EFPL": [
    {
      "lat": 60.9927778,
      "lon": 25.57,
      "name": "PÄIJÄT-HÄMEEN KESKUSSAIRAALA"
    }
  ],
  "EFPN": [
    {
      "lat": 61.728889,
      "lon": 29.393611,
      "name": "PUNKAHARJU"
    }
  ],
  "EFPO": [
    {
      "lat": 61.461389,
      "lon": 21.797778,
      "name": "PORI"
    }
  ],
  "EFPR": [
    {
      "lat": 60.479167,
      "lon": 26.593889,
      "name": "REDSTONE AERO"
    }
  ],
  "EFPT": [
    {
      "lat": 61.5047222,
      "lon": 23.8136111,
      "name": "TAMPEREEN YLIOPISTOLLINEN KESKUSSAIRAALA"
    }
  ],
  "EFPU": [
    {
      "lat": 65.402222,
      "lon": 26.946944,
      "name": "PUDASJARVI"
    }
  ],
  "EFPY": [
    {
      "lat": 63.729167,
      "lon": 25.931944,
      "name": "PYHASALMI"
    }
  ],
  "EFRA": [
    {
      "lat": 63.424167,
      "lon": 28.124167,
      "name": "RAUTAVAARA"
    }
  ],
  "EFRH": [
    {
      "lat": 64.688056,
      "lon": 24.695833,
      "name": "RAAHE-PATTIJOKI"
    }
  ],
  "EFRN": [
    {
      "lat": 62.065556,
      "lon": 28.356667,
      "name": "RANTASALMI"
    }
  ],
  "EFRO": [
    {
      "lat": 66.561667,
      "lon": 25.830833,
      "name": "ROVANIEMI"
    }
  ],
  "EFRU": [
    {
      "lat": 65.973056,
      "lon": 26.365278,
      "name": "RANUA"
    }
  ],
  "EFRV": [
    {
      "lat": 63.705556,
      "lon": 26.616389,
      "name": "KIURUVESI"
    }
  ],
  "EFRY": [
    {
      "lat": 60.744722,
      "lon": 24.107778,
      "name": "RAYSKALA"
    }
  ],
  "EFSA": [
    {
      "lat": 61.942778,
      "lon": 28.945,
      "name": "SAVONLINNA"
    }
  ],
  "EFSE": [
    {
      "lat": 61.064167,
      "lon": 26.795556,
      "name": "SELANPAA"
    }
  ],
  "EFSI": [
    {
      "lat": 62.693611,
      "lon": 22.831944,
      "name": "SEINAJOKI"
    }
  ],
  "EFSO": [
    {
      "lat": 67.396667,
      "lon": 26.618056,
      "name": "SODANKYLA"
    }
  ],
  "EFSU": [
    {
      "lat": 64.821944,
      "lon": 28.710278,
      "name": "SUOMUSSALMI"
    }
  ],
  "EFTO": [
    {
      "lat": 60.079167,
      "lon": 24.172222,
      "name": "TORBACKA"
    }
  ],
  "EFTP": [
    {
      "lat": 61.415278,
      "lon": 23.587778,
      "name": "TAMPERE/PIRKKALA"
    }
  ],
  "EFTS": [
    {
      "lat": 61.773056,
      "lon": 24.025278,
      "name": "TEISKO"
    }
  ],
  "EFTU": [
    {
      "lat": 60.514722,
      "lon": 22.261667,
      "name": "TURKU"
    }
  ],
  "EFTV": [
    {
      "lat": 60.4530556,
      "lon": 22.2958333,
      "name": "TURUN YLIOPISTOLLINEN KESKUSSAIRAALA"
    }
  ],
  "EFUT": [
    {
      "lat": 60.896389,
      "lon": 26.938056,
      "name": "UTTI"
    }
  ],
  "EFVA": [
    {
      "lat": 63.045278,
      "lon": 21.764167,
      "name": "VAASA"
    }
  ],
  "EFVI": [
    {
      "lat": 63.1225,
      "lon": 25.816111,
      "name": "VIITASAARI"
    }
  ],
  "EFVL": [
    {
      "lat": 64.501944,
      "lon": 26.76,
      "name": "VAALA"
    }
  ],
  "EFVP": [
    {
      "lat": 61.039722,
      "lon": 22.591667,
      "name": "VAMPULA"
    }
  ],
  "EFVT": [
    {
      "lat": 63.397778,
      "lon": 24.030556,
      "name": "SULKAHARJU"
    }
  ],
  "EFWB": [
    {
      "lat": 60.663611,
      "lon": 26.745833,
      "name": "WREDEBY"
    }
  ],
  "EFYL": [
    {
      "lat": 64.054722,
      "lon": 24.725278,
      "name": "YLIVIESKA"
    }
  ]
}`
