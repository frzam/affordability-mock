package controllers

type AffordabilityRequest struct {
	Accept            bool `json:"accept"`
	AffordabilityInfo struct {
		IDType                            int    `json:"idType"`
		IDNumber                          string `json:"idNumber"`
		Amount                            int    `json:"amount"`
		FirstName                         string `json:"firstName"`
		ThirdName                         string `json:"thirdName"`
		FamilyName                        string `json:"familyName"`
		Nationality                       int    `json:"nationality"`
		CityOfResidence                   int    `json:"cityOfResidence"`
		MaritalStatus                     int    `json:"maritalStatus"`
		NumberOfHouseholdhelp             int    `json:"numberOfHouseholdhelp"`
		HomeOwnership                     int    `json:"homeOwnership"`
		ResidentialType                   int    `json:"residentialType"`
		NumberOfDependents                int    `json:"numberOfDependents"`
		NumberOfDependentsInPrivateSchool int    `json:"numberOfDependentsInPrivateSchool"`
		NumberOfDependentsInPublicSchool  int    `json:"numberOfDependentsInPublicSchool"`
		EnquiryType                       int    `json:"enquiryType"`
		ProductType                       int    `json:"productType"`
	} `json:"affordabilityInfo"`
	IncomeComputationMetrics struct {
		BasicIncome      int64 `json:"basicIncome"`
		AdditionalIncome int64 `json:"additionalIncome"`
		IsRetired        bool  `json:"isRetired"`
		Breadwinner      bool  `json:"breadwinner"`
		IsREDFCustomer   bool  `json:"isREDFCustomer"`
	} `json:"incomeComputationMetrics"`
	CustomerDeclarationOfExpenses []struct {
		TypeID    int    `json:"typeId"`
		Code      string `json:"code"`
		Value     int    `json:"value"`
		IsVerfied bool   `json:"isVerfied"`
	} `json:"customerDeclarationOfExpenses"`
	ReferenceNumber string `json:"referenceNumber"`
	ResponseType    int    `json:"responseType"`
	RequestType     int    `json:"requestType"`
}
