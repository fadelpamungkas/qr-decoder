package helper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/qr-decoder/models"
)

var (
	pointOfInitiationMethod = map[string]string{
		"11": "Static QR",
		"12": "Dynamic QR",
	}

	merchantCriteria = map[string]string{
		"UMI": "Usaha Mikro",
		"UKE": "Usaha Kecil",
		"UME": "Usaha Menengah",
		"UBE": "Usaha Besar",
	}

	currencies = map[string]string{
		"360": "IDR|Rupiah",
		"840": "USD|US Dollar",
		"702": "SGD|Singapore Dollar",
		"764": "THB|Thai Baht",
		"458": "MYR|Ringgit",
		"608": "PHP|Peso",
		"104": "MMK|Kyat",
		"704": "VND|Dong",
		"156": "CNY|Yuan",
		"410": "KRW|Won",
		"392": "JPY|Yen",
	}

	tipsIndicator = map[string]string{
		"01": "User Tips",
		"02": "Fixed Tips",
		"03": "Percentage Tips",
	}

	merchantCategoryCodes = map[string]string{
		"0742": "Veterinary services",
		"0743": "Wine producers",
		"0744": "Champagne producers",
		"0763": "Agricultural Cooperatives",
		"0780": "Landscaping and horticultural services",
		"1353": "Dia (Spain)-Hypermarkets of Food",
		"1406": "H&M Moda (Spain)-Retail Merchants",
		"1520": "General contractors — residential and commercial",
		"1711": "Heating, plumbing and air-conditioning contractors",
		"1731": "Electrical contractors",
		"1740": "Masonry, stonework, tile setting, plastering and insulation contractors",
		"1750": "Carpentry contractors",
		"1761": "Roofing, siding and sheet metal work contractors",
		"1771": "Concrete work contractors",
		"1799": "Special trade contractors — not elsewhere classified",
		"2741": "Miscellaneous publishing and printing services",
		"2791": "Typesetting, platemaking and related services",
		"2842": "Speciality cleaning, polishing and sanitation preparations",
		"G300": "Airlines (codes between 3000 and 3350)",
		"G335": "Car rentals (codes between 3351 and 3500)",
		"G350": "Hotels (codes between 3501 and 3999)",
		"4011": "Railroads",
		"4111": "Local and suburban commuter passenger transportation, including ferries",
		"4112": "Passenger railways",
		"4119": "Ambulance Services",
		"4121": "Taxi-cabs and limousines",
		"4131": "Bus Lines",
		"4214": "Motor freight carriers and trucking — local and long distance, moving and storage companies and local delivery",
		"4215": "Courier services — air and ground and freight forwarders",
		"4225": "Public warehousing and storage — farm products, refrigerated goods and household goods",
		"4411": "Steamships and cruise lines",
		"4457": "Boat Rentals and Leasing",
		"4468": "Marinas, marine service and supplies",
		"4511": "Airlines and Air Carriers (Not Elsewhere Classified)",
		"4582": "Airports, Flying Fields, and Airport Terminals",
		"4722": "Travel agencies and tour operators",
		"4723": "Package Tour Operators – Germany Only",
		"4784": "Tolls and bridge fees",
		"4789": "Transportation services — not elsewhere classified",
		"4812": "Telecommunication equipment and telephone sales",
		"4813": "Key-entry Telecom Merchant providing single local and long-distance phone calls using a central access number in a non–face-to-face environment using key entry",
		"4814": "Telecommunication services, including local and long distance calls, credit card calls, calls through use of magnetic stripe reading telephones and faxes",
		"4815": "Monthly summary telephone charges",
		"4816": "Computer network/information services",
		"4821": "Telegraph services",
		"4829": "Wire transfers and money orders",
		"4899": "Cable and other pay television services",
		"4900": "Utilities — electric, gas, water and sanitary",
		"5013": "Motor vehicle supplies and new parts",
		"5021": "Office and commercial furniture",
		"5039": "Construction materials — not elsewhere classified",
		"5044": "Office, photographic, photocopy and microfilm equipment",
		"5045": "Computers, computer peripheral equipment — not elsewhere classified",
		"5046": "Commercial equipment — not elsewhere classified",
		"5047": "Dental/laboratory/medical/ophthalmic hospital equipment and supplies",
		"5051": "Metal service centres and offices",
		"5065": "Electrical parts and equipment",
		"5072": "Hardware equipment and supplies",
		"5074": "Plumbing and heating equipment and supplies",
		"5085": "Industrial supplies — not elsewhere classified",
		"5094": "Precious stones and metals, watches and jewellery",
		"5099": "Durable goods — not elsewhere classified",
		"5111": "Stationery, office supplies and printing and writing paper",
		"5122": "Drugs, drug proprietors",
		"5131": "Piece goods, notions and other dry goods",
		"5137": "Men’s, women’s and children’s uniforms and commercial clothing",
		"5139": "Commercial footwear",
		"5169": "Chemicals and allied products — not elsewhere classified",
		"5172": "Petroleum and petroleum products",
		"5192": "Books, Periodicals and Newspapers",
		"5193": "Florists’ supplies, nursery stock and flowers",
		"5198": "Paints, varnishes and supplies",
		"5199": "Non-durable goods — not elsewhere classified",
		"5200": "Home supply warehouse outlets",
		"5211": "Lumber and building materials outlets",
		"5231": "Glass, paint and wallpaper shops",
		"5251": "Hardware shops",
		"5261": "Lawn and garden supply outlets, including nurseries",
		"5262": "Marketplaces (online Marketplaces)",
		"5271": "Mobile home dealers",
		"5299": "Warehouse Club Gas",
		"5300": "Wholesale clubs",
		"5309": "Duty-free shops",
		"5310": "Discount shops",
		"5311": "Department stores",
		"5331": "Variety stores",
		"5333": "HYPERMARKETS OF FOOD",
		"5399": "Miscellaneous general merchandise",
		"5411": "Groceries and supermarkets",
		"5422": "Freezer and locker meat provisioners",
		"5441": "Candy, nut and confectionery shops",
		"5451": "Dairies",
		"5462": "Bakeries",
		"5499": "Miscellaneous food shops — convenience and speciality retail outlets",
		"5511": "Car and truck dealers (new and used) sales, services, repairs, parts and leasing",
		"5521": "Car and truck dealers (used only) sales, service, repairs, parts and leasing",
		"5531": "Auto Store",
		"5532": "Automotive Tire Stores",
		"5533": "Automotive Parts and Accessories Stores",
		"5541": "Service stations (with or without ancillary services)",
		"5542": "Automated Fuel Dispensers",
		"5551": "Boat Dealers",
		"5552": "Electric Vehicle Charging",
		"5561": "Camper, recreational and utility trailer dealers",
		"5571": "Motorcycle shops and dealers",
		"5592": "Motor home dealers",
		"5598": "Snowmobile dealers",
		"5599": "Miscellaneous automotive, aircraft and farm equipment dealers — not elsewhere classified",
		"5611": "Men’s and boys’ clothing and accessory shops",
		"5621": "Women’s ready-to-wear shops",
		"5631": "Women’s accessory and speciality shops",
		"5641": "Children’s and infants’ wear shops",
		"5651": "Family clothing shops",
		"5655": "Sports and riding apparel shops",
		"5661": "Shoe shops",
		"5681": "Furriers and fur shops",
		"5691": "Men’s and women’s clothing shops",
		"5697": "Tailors, seamstresses, mending and alterations",
		"5698": "Wig and toupee shops",
		"5699": "Miscellaneous apparel and accessory shops",
		"5712": "Furniture, home furnishings and equipment shops and manufacturers, except appliances",
		"5713": "Floor covering services",
		"5714": "Drapery, window covering and upholstery shops",
		"5715": "Alcoholic beverage wholesalers",
		"5718": "Fireplaces, fireplace screens and accessories shops",
		"5719": "Miscellaneous home furnishing speciality shops",
		"5722": "Household appliance shops",
		"5732": "Electronics shops",
		"5733": "Music shops — musical instruments, pianos and sheet music",
		"5734": "Computer software outlets",
		"5735": "Record shops",
		"5811": "Caterers",
		"5812": "Eating places and restaurants",
		"5813": "Drinking places (alcoholic beverages) — bars, taverns, night-clubs, cocktail lounges and discothèques",
		"5814": "Fast food restaurants",
		"5815": "Digital Goods Media – Books, Movies, Music",
		"5816": "Digital Goods – Games",
		"5817": "Digital Goods – Applications (Excludes Games)",
		"5818": "Digital Goods – Large Digital Goods Merchant",
		"5912": "Drug stores and pharmacies",
		"5921": "Package shops — beer, wine and liquor",
		"5931": "Used merchandise and second-hand shops",
		"5932": "Antique Shops – Sales, Repairs, and Restoration Services",
		"5933": "Pawn shops",
		"5935": "Wrecking and salvage yards",
		"5937": "Antique Reproductions",
		"5940": "Bicycle Shops – Sales and Service",
		"5941": "Sporting goods shops",
		"5942": "Book Stores",
		"5943": "Stationery, office and school supply shops",
		"5944": "Jewellery, watch, clock and silverware shops",
		"5945": "Hobby, toy and game shops",
		"5946": "Camera and photographic supply shops",
		"5947": "Gift, card, novelty and souvenir shops",
		"5948": "Luggage and leather goods shops",
		"5949": "Sewing, needlework, fabric and piece goods shops",
		"5950": "Glassware and crystal shops",
		"5960": "Direct marketing — insurance services",
		"5961": "Mail Order Houses Including Catalog Order Stores",
		"5962": "Telemarketing — travel-related arrangement services",
		"5963": "Door-to-door sales",
		"5964": "Direct marketing — catalogue merchants",
		"5965": "Direct marketing — combination catalogue and retail merchants",
		"5966": "Direct marketing — outbound telemarketing merchants",
		"5967": "Direct marketing — inbound telemarketing merchants",
		"5968": "Direct marketing — continuity/subscription merchants",
		"5969": "Direct marketing/direct marketers — not elsewhere classified",
		"5970": "Artist’s Supply and Craft Shops",
		"5971": "Art Dealers and Galleries",
		"5972": "Stamp and coin shops",
		"5973": "Religious goods and shops",
		"5974": "Rubber Stamp Store",
		"5975": "Hearing aids — sales, service and supplies",
		"5976": "Orthopaedic goods and prosthetic devices",
		"5977": "Cosmetic Stores",
		"5978": "Typewriter outlets — sales, service and rentals",
		"5983": "Fuel dealers — fuel oil, wood, coal and liquefied petroleum",
		"5992": "Florists",
		"5993": "Cigar shops and stands",
		"5994": "Newsagents and news-stands",
		"5995": "Pet shops, pet food and supplies",
		"5996": "Swimming pools — sales, supplies and services",
		"5997": "Electric razor outlets — sales and service",
		"5998": "Tent and awning shops",
		"5999": "Miscellaneous and speciality retail outlets",
		"6010": "Financial institutions — manual cash disbursements",
		"6011": "Financial institutions — automated cash disbursements",
		"6012": "Financial institutions — merchandise and services",
		"6050": "Quasi Cash—Customer Financial Institution",
		"6051": "Non-financial institutions — foreign currency, money orders (not wire transfer), scrip and travellers’ checks",
		"6211": "Securities — brokers and dealers",
		"6300": "Insurance sales, underwriting and premiums",
		"6381": "Insurance–Premiums",
		"6399": "Insurance, Not Elsewhere Classified ( no longer valid for first presentment work)",
		"6513": "Real Estate Agents and Managers",
		"6529": "Remote Stored Value Load — Member Financial Institution",
		"6530": "Remove Stored Value Load — Merchant",
		"6532": "Payment Transaction—Customer Financial Institution",
		"6533": "Payment Transaction—Merchant",
		"6535": "Value Purchase–Member Financial Institution",
		"6536": "MoneySend Intracountry",
		"6537": "MoneySend Intercountry",
		"6538": "MoneySend Funding",
		"6539": "Funding Transaction (Excluding MoneySend)",
		"6540": "Non-Financial Institutions – Stored Value Card Purchase/Load",
		"6611": "Overpayments",
		"6760": "Savings Bonds",
		"7011": "Lodging — hotels, motels and resorts",
		"7012": "Timeshares",
		"7032": "Sporting and recreational camps",
		"7033": "Trailer parks and camp-sites",
		"7210": "Laundry, cleaning and garment services",
		"7211": "Laundry services — family and commercial",
		"7216": "Dry cleaners",
		"7217": "Carpet and upholstery cleaning",
		"7221": "Photographic studios",
		"7230": "Barber and Beauty Shops",
		"7251": "Shoe repair shops, shoe shine parlours and hat cleaning shops",
		"7261": "Funeral services and crematoriums",
		"7273": "Dating and escort services",
		"7276": "Tax preparation services",
		"7277": "Counselling services — debt, marriage and personal",
		"7278": "Buying and shopping services and clubs",
		"7280": "Hospital Patient-Personal Funds Withdrawal",
		"7295": "Babysitting Services",
		"7296": "Clothing rentals — costumes, uniforms and formal wear",
		"7297": "Massage parlours",
		"7298": "Health and beauty spas",
		"7299": "Miscellaneous personal services — not elsewhere classified",
		"7311": "Advertising Services",
		"7321": "Consumer credit reporting agencies",
		"7322": "Debt collection agencies",
		"7332": "Blueprinting and Photocopying Services",
		"7333": "Commercial photography, art and graphics",
		"7338": "Quick copy, reproduction and blueprinting services",
		"7339": "Stenographic and secretarial support services",
		"7342": "Exterminating and disinfecting services",
		"7349": "Cleaning, maintenance and janitorial services",
		"7361": "Employment agencies and temporary help services",
		"7372": "Computer programming, data processing and integrated systems design services",
		"7375": "Information retrieval services",
		"7379": "Computer maintenance and repair services — not elsewhere classified",
		"7392": "Management, consulting and public relations services",
		"7393": "Detective  agencies,  protective  agencies  and  security  services,  including  armoured  cars  and guard dogs",
		"7394": "Equipment, tool, furniture and appliance rentals and leasing",
		"7395": "Photofinishing laboratories and photo developing",
		"7399": "Business services — not elsewhere classified",
		"7511": "Truck Stop",
		"7512": "Automobile Rental Agency—not elsewhere classified",
		"7513": "Truck and utility trailer rentals",
		"7519": "Motor home and recreational vehicle rentals",
		"7523": "Parking lots and garages",
		"7524": "Express Payment Service Merchants–Parking Lots and Garages",
		"7531": "Automotive Body Repair Shops",
		"7534": "Tyre retreading and repair shops",
		"7535": "Automotive Paint Shops",
		"7538": "Automotive Service Shops (Non-Dealer)",
		"7539": "Automotive Service Shops (Spain) - Other Merchant Categories",
		"7542": "Car washes",
		"7549": "Towing services",
		"7622": "Electronics repair shops",
		"7623": "Air Conditioning and Refrigeration Repair Shops",
		"7629": "Electrical and small appliance repair shops",
		"7631": "Watch, clock and jewellery repair shops",
		"7641": "Furniture reupholstery, repair and refinishing",
		"7692": "Welding services",
		"7699": "Miscellaneous repair shops and related services",
		"7800": "Government-Owned Lotteries (US Region only)",
		"7801": "Government Licensed On-Line Casinos (On-Line Gambling) (US Region only)",
		"7802": "Government-Licensed Horse/Dog Racing (US Region only)",
		"7829": "Motion picture and video tape production and distribution",
		"7832": "Motion picture theatres",
		"7833": "Express Payment Service — Motion Picture Theater",
		"7841": "Video tape rentals",
		"7911": "Dance halls, studios and schools",
		"7922": "Theatrical producers (except motion pictures) and ticket agencies",
		"7929": "Bands, Orchestras, and Miscellaneous Entertainers (Not Elsewhere Classified)",
		"7932": "Billiard and Pool Establishments",
		"7933": "Bowling Alleys",
		"7941": "Commercial sports, professional sports clubs, athletic fields and sports promoters",
		"7991": "Tourist attractions and exhibits",
		"7992": "Public golf courses",
		"7993": "Video amusement game supplies",
		"7994": "Video game arcades and establishments",
		"7995": "Betting, including Lottery Tickets, Casino Gaming Chips, Off-Track Betting, and Wagers at Race Tracks",
		"7996": "Amusement Parks, Circuses, Carnivals, and Fortune Tellers",
		"7997": "Membership clubs (sports, recreation, athletic), country clubs and private golf courses",
		"7998": "Aquariums, Seaquariums, Dolphinariums, and Zoos",
		"7999": "Recreation services — not elsewhere classified",
		"8011": "Doctors and physicians — not elsewhere classified",
		"8021": "Dentists and orthodontists",
		"8031": "Osteopaths",
		"8041": "Chiropractors",
		"8042": "Optometrists and ophthalmologists",
		"8043": "Opticians, optical goods and eyeglasses",
		"8044": "Optical Goods and Eyeglasses",
		"8049": "Podiatrists and chiropodists",
		"8050": "Nursing and personal care facilities",
		"8062": "Hospitals",
		"8071": "Medical and dental laboratories",
		"8099": "Medical services and health practitioners — not elsewhere classified",
		"8111": "Legal services and attorneys",
		"8211": "Elementary and secondary schools",
		"8220": "Colleges, universities, professional schools and junior colleges",
		"8241": "Correspondence schools",
		"8244": "Business and secretarial schools",
		"8249": "Trade and vocational schools",
		"8299": "Schools and educational services — not elsewhere classified",
		"8351": "Child care services",
		"8398": "Charitable and social service organizations",
		"8641": "Civic, social and fraternal associations",
		"8651": "Political organizations",
		"8661": "Religious organizations",
		"8675": "Automobile Associations",
		"8699": "Membership organization — not elsewhere classified",
		"8734": "Testing laboratories (non-medical)",
		"8911": "Architectural, Engineering, and Surveying Services",
		"8931": "Accounting, Auditing, and Bookkeeping Services",
		"8999": "Professional services — not elsewhere classified",
		"9034": "I-Purchasing Pilot",
		"9211": "Court costs, including alimony and child support",
		"9222": "Fines",
		"9223": "Bail and Bond Payments",
		"9311": "Tax payments",
		"9399": "Government services — not elsewhere classified",
		"9402": "Postal services — government only",
		"9405": "U.S. Federal Government Agencies or Departments",
		"9406": "Government-Owned Lotteries (Non-U.S. region)",
		"9700": "Automated Referral Service",
		"9701": "Visa Credential Server",
		"9702": "Emergency Services (GCAS) (Visa use only)",
		"9751": "UK Supermarkets, Electronic Hot File",
		"9752": "UK Petrol Stations, Electronic Hot File",
		"9754": "Gambling-Horse, Dog Racing, State Lottery",
		"9950": "Intra-Company Purchases",
	}
)

func ParseQRCode(data string) models.QRData {
	qrData := models.QRData{
		RawData:             data,
		MerchantAccountInfo: make(map[string]models.MerchantAccountInfo),
	}

	for len(data) > 0 {
		tag := data[:2]
		length, _ := strconv.Atoi(data[2:4])
		value := data[4 : 4+length]
		data = data[4+length:]

		tagNum, _ := strconv.Atoi(tag)

		switch {
		case tagNum == 0:
			qrData.PayloadFormatIndicator = value
		case tagNum == 1:
			qrData.PointOfInitiationMethod = GetPointOfInitiationMethod(value)
		case tagNum >= 2 && tagNum <= 51:
			qrData.MerchantAccountInfo[tag] = ParseMerchantAccountInfo(value)
		case tagNum == 52:
			qrData.MerchantCategoryCode = GetMerchantCategoryCode(value)
		case tagNum == 53:
			qrData.TransactionCurrency = GetTransactionCurrency(value)
		case tagNum == 54:
			qrData.TransactionAmount = value
		case tagNum == 55:
			qrData.TipOrConvenienceIndicator = GetTipsIndicator(value)
		case tagNum == 56:
			qrData.ValueOfConvenienceFeeFixed = value
		case tagNum == 57:
			qrData.ValueOfConvenienceFeePercentage = value
		case tagNum == 58:
			qrData.CountryCode = value
		case tagNum == 59:
			qrData.MerchantName = value
		case tagNum == 60:
			qrData.MerchantCity = value
		case tagNum == 61:
			qrData.PostalCode = value
		case tagNum == 62:
			qrData.AdditionalDataFieldTemplate = value
		case tagNum == 63:
			qrData.CRC = value
		}
	}

	return qrData
}

func ParseMerchantAccountInfo(data string) models.MerchantAccountInfo {
	mai := models.MerchantAccountInfo{}
	for len(data) > 0 {
		subTag := data[:2]
		length, _ := strconv.Atoi(data[2:4])
		value := data[4 : 4+length]
		data = data[4+length:]

		switch subTag {
		case "00":
			mai.GlobalUniqueIdentifier = reverseGlobalUniqueIdentifier(value)
		case "01":
			mai.MerchantPAN = value
		case "02":
			mai.MerchantID = value
		case "03":
			mai.MerchantCriteria = GetMerchantCriteria(value)
		}
	}
	return mai
}

func reverseGlobalUniqueIdentifier(data string) string {
	if len(data) == 0 {
		return data
	}

	splitData := strings.Split(data, ".")
	for i := 0; i < len(splitData)/2; i++ {
		j := len(splitData) - i - 1
		splitData[i], splitData[j] = splitData[j], splitData[i]
	}

	return strings.Join(splitData, ".")
}

func GetPointOfInitiationMethod(poi string) string {
	if poi == "" {
		return poi
	}

	if value, ok := pointOfInitiationMethod[poi]; ok {
		return fmt.Sprintf("%s (%s)", poi, value)
	}

	return poi
}

func PrintToJSON(data models.QRData) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	return nil
}

func GetTransactionCurrency(currency string) string {
	if currency == "" {
		return currency
	}

	if value, ok := currencies[currency]; ok {
		return fmt.Sprintf("%s (%s)", currency, value)
	}

	return currency
}

func GetTipsIndicator(value string) string {
	if value == "" {
		return value
	}

	if value, ok := tipsIndicator[value]; ok {
		return fmt.Sprintf("%s (%s)", value, value)
	}

	return value

}

func GetMerchantCriteria(criteria string) string {
	if criteria == "" {
		return criteria
	}

	if value, ok := merchantCriteria[criteria]; ok {
		return fmt.Sprintf("%s (%s)", criteria, value)
	}

	return criteria
}

func GetMerchantCategoryCode(code string) string {
	if code == "" {
		return code
	}

	if value, ok := merchantCategoryCodes[code]; ok {
		return fmt.Sprintf("%s (%s)", code, value)
	}

	return code
}
