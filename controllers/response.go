package controllers

import "time"

type AffordabilityResponse struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
	Data      struct {
		ReportDate    string `json:"reportDate"`
		ReportDetails struct {
			ReportDate         string `json:"reportDate"`
			EnquiryType        string `json:"enquiryType"`
			ProductType        string `json:"productType"`
			EnquiryNumber      string `json:"enquiryNumber"`
			NumberOfApplicants string `json:"numberOfApplicants"`
			AccountType        string `json:"accountType"`
			ReferenceNumber    string `json:"referenceNumber"`
			Amount             string `json:"amount"`
			MemberType         struct {
				ID     int    `json:"id"`
				Code   string `json:"code"`
				NameAr string `json:"nameAr"`
				Name   string `json:"name"`
			} `json:"memberType"`
			Status struct {
				ID     int    `json:"id"`
				Code   string `json:"code"`
				NameAr string `json:"nameAr"`
				Name   string `json:"name"`
			} `json:"status"`
		} `json:"reportDetails"`
		ProvidedDemographicsInfo struct {
			DemIDNumber string `json:"demIDNumber"`
			DemIDType   struct {
				TypeID     int    `json:"typeID"`
				TypeNameEN string `json:"typeNameEN"`
				TypeNameAR string `json:"typeNameAR"`
				IDTypeCode string `json:"idTypeCode"`
			} `json:"demIDType"`
			DemIDExpiryDate  string `json:"demIDExpiryDate"`
			DemApplicantType struct {
				ApplicantTypeID     int    `json:"applicantTypeID"`
				ApplicantTypeCode   string `json:"applicantTypeCode"`
				ApplicantTypeNameEN string `json:"applicantTypeNameEN"`
				ApplicantTypeNameAR string `json:"applicantTypeNameAR"`
				IsActive            bool   `json:"isActive"`
			} `json:"demApplicantType"`
			DemCustomerName   string `json:"demCustomerName"`
			DemFamilyName     string `json:"demFamilyName"`
			DemFirstName      string `json:"demFirstName"`
			DemSecondName     string `json:"demSecondName"`
			DemThirdName      string `json:"demThirdName"`
			DemCustomerNameAr string `json:"demCustomerNameAr"`
			DemFamilyNameAr   string `json:"demFamilyNameAr"`
			DemFirstNameAr    string `json:"demFirstNameAr"`
			DemSecondNameAr   string `json:"demSecondNameAr"`
			DemThirdNameAr    string `json:"demThirdNameAr"`
			DemDateOfBirth    string `json:"demDateOfBirth"`
			DemGender         string `json:"demGender"`
			DemMaritalStatus  struct {
				MatrialStatusID   int    `json:"matrialStatusId"`
				StatusNameEN      string `json:"statusNameEN"`
				StatusNameAR      string `json:"statusNameAR"`
				MaritalStatusCode string `json:"maritalStatusCode"`
			} `json:"demMaritalStatus"`
			DemNationality struct {
				Couid     int    `json:"couid"`
				CouNameEN string `json:"couNameEN"`
				CouNameAR string `json:"couNameAR"`
				CouCode   string `json:"couCode"`
			} `json:"demNationality"`
		} `json:"providedDemographicsInfo"`
		AvailableDemographicsInfo struct {
			DemIDNumber string `json:"demIDNumber"`
			DemIDType   struct {
				TypeID     int    `json:"typeID"`
				TypeNameEN string `json:"typeNameEN"`
				TypeNameAR string `json:"typeNameAR"`
				IDTypeCode string `json:"idTypeCode"`
			} `json:"demIDType"`
			DemIDExpiryDate  string `json:"demIDExpiryDate"`
			DemApplicantType struct {
				ApplicantTypeID     int    `json:"applicantTypeID"`
				ApplicantTypeCode   string `json:"applicantTypeCode"`
				ApplicantTypeNameEN string `json:"applicantTypeNameEN"`
				ApplicantTypeNameAR string `json:"applicantTypeNameAR"`
				IsActive            bool   `json:"isActive"`
			} `json:"demApplicantType"`
			DemCustomerName   string `json:"demCustomerName"`
			DemFamilyName     string `json:"demFamilyName"`
			DemFirstName      string `json:"demFirstName"`
			DemSecondName     string `json:"demSecondName"`
			DemThirdName      string `json:"demThirdName"`
			DemCustomerNameAr string `json:"demCustomerNameAr"`
			DemFamilyNameAr   string `json:"demFamilyNameAr"`
			DemFirstNameAr    string `json:"demFirstNameAr"`
			DemSecondNameAr   string `json:"demSecondNameAr"`
			DemThirdNameAr    string `json:"demThirdNameAr"`
			DemDateOfBirth    string `json:"demDateOfBirth"`
			DemGender         string `json:"demGender"`
			DemMaritalStatus  struct {
				MatrialStatusID   int    `json:"matrialStatusId"`
				StatusNameEN      string `json:"statusNameEN"`
				StatusNameAR      string `json:"statusNameAR"`
				MaritalStatusCode string `json:"maritalStatusCode"`
			} `json:"demMaritalStatus"`
			DemNationality struct {
				Couid     int    `json:"couid"`
				CouNameEN string `json:"couNameEN"`
				CouNameAR string `json:"couNameAR"`
				CouCode   string `json:"couCode"`
			} `json:"demNationality"`
		} `json:"availableDemographicsInfo"`
		PrevEnquiries []struct {
			PrevEnqDate string `json:"prevEnqDate"`
			PreEnqType  struct {
				EnqTypeCode          string `json:"enqTypeCode"`
				EnqTypeDescriptionAr string `json:"enqTypeDescriptionAr"`
				EnqTypeDescriptionEn string `json:"enqTypeDescriptionEn"`
			} `json:"preEnqType"`
			PrevEnqEnquirer struct {
				MemberCode   string `json:"memberCode"`
				MemberNameEN string `json:"memberNameEN"`
				MemberNameAR string `json:"memberNameAR"`
			} `json:"prevEnqEnquirer"`
			PrevEnqMemberRef       string `json:"prevEnqMemberRef"`
			PrevEnquirerName       string `json:"prevEnquirerName"`
			PrevEnquirerNameAr     string `json:"prevEnquirerNameAr"`
			PrevEnqProductTypeDesc struct {
				ProductID int    `json:"productId"`
				Code      string `json:"code"`
				TextEn    string `json:"textEn"`
				TextAr    string `json:"textAr"`
			} `json:"prevEnqProductTypeDesc"`
			PrevEnqAmount int `json:"prevEnqAmount"`
			OtherReason   struct {
				EnquiryReasonCodeDescEn string `json:"enquiryReasonCodeDescEn"`
				EnquiryReasonCodeDescAr string `json:"enquiryReasonCodeDescAr"`
				EnquiryReasonCodeName   string `json:"enquiryReasonCodeName"`
			} `json:"otherReason"`
		} `json:"prevEnquiries"`
		CreditInstrumentDetails []struct {
			CiCreditor struct {
				MemberCode   string `json:"memberCode"`
				MemberNameEN string `json:"memberNameEN"`
				MemberNameAR string `json:"memberNameAR"`
			} `json:"ciCreditor"`
			CiProductTypeDesc struct {
				ProductID int    `json:"productId"`
				Code      string `json:"code"`
				TextEn    string `json:"textEn"`
				TextAr    string `json:"textAr"`
			} `json:"ciProductTypeDesc"`
			CiAccountNumber  string `json:"ciAccountNumber"`
			CiLimit          int    `json:"ciLimit"`
			CiIssuedDate     string `json:"ciIssuedDate"`
			CiExpirationDate string `json:"ciExpirationDate"`
			CiStatus         struct {
				CreditInstrumentStatusCode   string `json:"creditInstrumentStatusCode"`
				CreditInstrumentStatusDescAr string `json:"creditInstrumentStatusDescAr"`
				CreditInstrumentStatusDescEn string `json:"creditInstrumentStatusDescEn"`
			} `json:"ciStatus"`
			CiClosingDate      string `json:"ciClosingDate"`
			CiTenure           int    `json:"ciTenure"`
			CiPaymentFrequency struct {
				PaymentFrequencyCodeDescEn string `json:"paymentFrequencyCodeDescEn"`
				PaymentFrequencyCodeDescAr string `json:"paymentFrequencyCodeDescAr"`
				PaymentFrequencyCodeName   string `json:"paymentFrequencyCodeName"`
			} `json:"ciPaymentFrequency"`
			CiInstallmentAmount    int `json:"ciInstallmentAmount"`
			CiSalaryAssignmentFlag struct {
				SalaryAssignmentFlagDescEn string `json:"salaryAssignmentFlagDescEn"`
				SalaryAssignmentFlagDescAr string `json:"salaryAssignmentFlagDescAr"`
				SalaryAssignmentFlagCode   string `json:"salaryAssignmentFlagCode"`
			} `json:"ciSalaryAssignmentFlag"`
			CiConsumerSecurityType struct {
				ConsumerSecurityTypeDescEn string `json:"consumerSecurityTypeDescEn"`
				ConsumerSecurityTypeDescAr string `json:"consumerSecurityTypeDescAr"`
				ConsumerSecurityTypeCode   string `json:"consumerSecurityTypeCode"`
			} `json:"ciConsumerSecurityType"`
			CiOutstandingBalance  int    `json:"ciOutstandingBalance"`
			CiPastDue             int    `json:"ciPastDue"`
			CiLastAmountPaid      int    `json:"ciLastAmountPaid"`
			CiLastPaymentDate     string `json:"ciLastPaymentDate"`
			CiAsOfDate            string `json:"ciAsOfDate"`
			CiNextDueDate         string `json:"ciNextDueDate"`
			CiSummary             string `json:"ciSummary"`
			CiBalloonPayment      string `json:"ciBalloonPayment"`
			CiDownPayment         int    `json:"ciDownPayment"`
			CiDispensedAmount     int    `json:"ciDispensedAmount"`
			CiMaxInstalmentAmount int    `json:"ciMaxInstalmentAmount"`
			CiSubProduct          struct {
				NameEn string `json:"nameEn"`
				NameAr string `json:"nameAr"`
				Code   string `json:"code"`
			} `json:"ciSubProduct"`
			MultiInstalmentDetails []struct {
				StartDate        time.Time `json:"startDate"`
				InstalmentAmount int       `json:"instalmentAmount"`
			} `json:"multiInstalmentDetails"`
			JointApplicantDetail struct {
				NumberOfApplicants          int    `json:"numberOfApplicants"`
				PercentageAllocation        int    `json:"percentageAllocation"`
				ApplicationInstalmentAmount int    `json:"applicationInstalmentAmount"`
				ApplicantLimit              int    `json:"applicantLimit"`
				LastAmountPaid              int    `json:"lastAmountPaid"`
				PaymentStatusEn             string `json:"paymentStatusEn"`
				PaymentStatusAr             string `json:"paymentStatusAr"`
				PaymentStatusCode           string `json:"paymentStatusCode"`
				LastPaymentDate             string `json:"lastPaymentDate"`
				PastDueAmount               int    `json:"pastDueAmount"`
				OutstandingBalance          int    `json:"outstandingBalance"`
				NextPaymentDate             string `json:"nextPaymentDate"`
			} `json:"jointApplicantDetail"`
			CiAverageInstallmentAmount struct {
			} `json:"ciAverageInstallmentAmount"`
			CiNumberOfApplicants string `json:"ciNumberOfApplicants"`
			JointApplicantFlag   struct {
				Code   string `json:"code"`
				TextEn string `json:"textEn"`
				TextAr string `json:"textAr"`
			} `json:"jointApplicantFlag"`
		} `json:"creditInstrumentDetails"`
		PrimaryDefaults []struct {
			PDefProductTypeDesc struct {
				ProductID int    `json:"productId"`
				Code      string `json:"code"`
				TextEn    string `json:"textEn"`
				TextAr    string `json:"textAr"`
			} `json:"pDefProductTypeDesc"`
			PDefAccountNo string `json:"pDefAccountNo"`
			PDefCreditor  struct {
				MemberCode   string `json:"memberCode"`
				MemberNameEN string `json:"memberNameEN"`
				MemberNameAR string `json:"memberNameAR"`
			} `json:"pDefCreditor"`
			PDefDateLoaded         string `json:"pDefDateLoaded"`
			PDefOriginalAmount     int    `json:"pDefOriginalAmount"`
			PDefOutstandingBalance int    `json:"pDefOutstandingBalance"`
			PDefaultStatuses       struct {
				DefaultStatusDescEn string `json:"defaultStatusDescEn"`
				DefaultStatusDescAr string `json:"defaultStatusDescAr"`
				DefaultStatusCode   string `json:"defaultStatusCode"`
			} `json:"pDefaultStatuses"`
			PDefSetteledDate string `json:"pDefSetteledDate"`
		} `json:"primaryDefaults"`
		GuarantorDefaults []struct {
			GDefProductTypeDesc struct {
				ProductID int    `json:"productId"`
				Code      string `json:"code"`
				TextEn    string `json:"textEn"`
				TextAr    string `json:"textAr"`
			} `json:"gDefProductTypeDesc"`
			GDefAccountNo string `json:"gDefAccountNo"`
			GDefCreditor  struct {
				MemberCode   string `json:"memberCode"`
				MemberNameEN string `json:"memberNameEN"`
				MemberNameAR string `json:"memberNameAR"`
			} `json:"gDefCreditor"`
			GDefDateLoaded         string `json:"gDefDateLoaded"`
			GDefOriginalAmount     int    `json:"gDefOriginalAmount"`
			GDefOutstandingBalance int    `json:"gDefOutstandingBalance"`
			GDefaultStatuses       struct {
				DefaultStatusDescEn string `json:"defaultStatusDescEn"`
				DefaultStatusDescAr string `json:"defaultStatusDescAr"`
				DefaultStatusCode   string `json:"defaultStatusCode"`
			} `json:"gDefaultStatuses"`
			GDefSetteledDate string `json:"gDefSetteledDate"`
		} `json:"guarantorDefaults"`
		BouncedCheques []struct {
			BcCheqDateLoaded  string `json:"bcCheqDateLoaded"`
			BcProductTypeDesc struct {
				ProductID int    `json:"productId"`
				Code      string `json:"code"`
				TextEn    string `json:"textEn"`
				TextAr    string `json:"textAr"`
			} `json:"bcProductTypeDesc"`
			BcCreditor struct {
				MemberCode   string `json:"memberCode"`
				MemberNameEN string `json:"memberNameEN"`
				MemberNameAR string `json:"memberNameAR"`
			} `json:"bcCreditor"`
			BcChequeNumber       string `json:"bcChequeNumber"`
			BcBalance            int    `json:"bcBalance"`
			BcOutstandingBalance int    `json:"bcOutstandingBalance"`
			BcDefaultStatuses    struct {
				DefaultStatusDescEn string `json:"defaultStatusDescEn"`
				DefaultStatusDescAr string `json:"defaultStatusDescAr"`
				DefaultStatusCode   string `json:"defaultStatusCode"`
			} `json:"bcDefaultStatuses"`
			BcSetteledDate string `json:"bcSetteledDate"`
		} `json:"bouncedCheques"`
		PublicNotices []struct {
			PublicNoDateLoaded string `json:"publicNoDateLoaded"`
			PublicNoTypes      struct {
				PublicNoticeTypeDescEn string `json:"publicNoticeTypeDescEn"`
				PublicNoticeTypeDescAr string `json:"publicNoticeTypeDescAr"`
				PublicNoticeTypeCode   string `json:"publicNoticeTypeCode"`
			} `json:"publicNoTypes"`
			PublicNoTextDescAr string `json:"publicNoTextDescAr"`
			PublicNoTextDescEn string `json:"publicNoTextDescEn"`
			PublicNoComment    string `json:"publicNoComment"`
			PublicNoCommentAr  string `json:"publicNoCommentAr"`
		} `json:"publicNotices"`
		Judgements []struct {
			ExecutionDate         string `json:"executionDate"`
			ResolutionNumber      string `json:"resolutionNumber"`
			CityNameEn            string `json:"cityNameEn"`
			CityNameAr            string `json:"cityNameAr"`
			CourtNameEn           string `json:"courtNameEn"`
			CourtNameAr           string `json:"courtNameAr"`
			LegalCaseNumber       string `json:"legalCaseNumber"`
			LoadedDate            string `json:"loadedDate"`
			OriginalClaimedAmount string `json:"originalClaimedAmount"`
			OutstandingBalance    string `json:"outstandingBalance"`
			SettlementDate        string `json:"settlementDate"`
			StatusNameEn          string `json:"statusNameEn"`
			StatusNameAr          string `json:"statusNameAr"`
			ExecutionType         string `json:"executionType"`
			StatusCode            string `json:"statusCode"`
			CityCode              string `json:"cityCode"`
		} `json:"judgements"`
		MemberNarratives []struct {
			NarrDateLoaded string `json:"narrDateLoaded"`
			NarrLoadedBy   struct {
				MemberCode   string `json:"memberCode"`
				MemberNameEN string `json:"memberNameEN"`
				MemberNameAR string `json:"memberNameAR"`
			} `json:"narrLoadedBy"`
			NarrativeTypes struct {
				NarrativeTypeDescAr string `json:"narrativeTypeDescAr"`
				NarrativeTypeDescEn string `json:"narrativeTypeDescEn"`
				NarrativeTypeCode   string `json:"narrativeTypeCode"`
			} `json:"narrativeTypes"`
			NarrTextDescAr string `json:"narrTextDescAr"`
			NarrTextDescEn string `json:"narrTextDescEn"`
		} `json:"memberNarratives"`
		PersonalNarratives []struct {
			NarrDateLoaded string `json:"narrDateLoaded"`
			NarrativeTypes struct {
				NarrativeTypeDescAr string `json:"narrativeTypeDescAr"`
				NarrativeTypeDescEn string `json:"narrativeTypeDescEn"`
				NarrativeTypeCode   string `json:"narrativeTypeCode"`
			} `json:"narrativeTypes"`
			NarrTextDescAr string `json:"narrTextDescAr"`
			NarrTextDescEn string `json:"narrTextDescEn"`
		} `json:"personalNarratives"`
		Addresses []struct {
			AdrsDateLoaded   string `json:"adrsDateLoaded"`
			AdrsAddressTypes struct {
				AddressID       int    `json:"addressID"`
				AddressTypeCode string `json:"addressTypeCode"`
				AddressNameAR   string `json:"addressNameAR"`
				AddressNameEN   string `json:"addressNameEN"`
			} `json:"adrsAddressTypes"`
			AdrsAddressLineFirstDescAr  string `json:"adrsAddressLineFirstDescAr"`
			AdrsAddressLineFirstDescEn  string `json:"adrsAddressLineFirstDescEn"`
			AdrsAddressLineSecondDescAr string `json:"adrsAddressLineSecondDescAr"`
			AdrsAddressLineSecondDescEn string `json:"adrsAddressLineSecondDescEn"`
			AdrsPOBox                   string `json:"adrsPOBox"`
			AdrsPostalCode              string `json:"adrsPostalCode"`
			AdrsCityDescAr              string `json:"adrsCityDescAr"`
			AdrsCityDescEn              string `json:"adrsCityDescEn"`
			NationalAddress             struct {
				BuildingNumber   string `json:"buildingNumber"`
				StreetAr         string `json:"streetAr"`
				StreetEn         string `json:"streetEn"`
				DistrictAr       string `json:"districtAr"`
				DistrictEn       string `json:"districtEn"`
				AdditionalNumber string `json:"additionalNumber"`
				UnitNumber       string `json:"unitNumber"`
			} `json:"nationalAddress"`
		} `json:"addresses"`
		Contacts []struct {
			ConNumberTypes struct {
				ContactNumberTypeCode          string `json:"contactNumberTypeCode"`
				ContactNumberTypeDescriptionAr string `json:"contactNumberTypeDescriptionAr"`
				ContactNumberTypeDescriptionEn string `json:"contactNumberTypeDescriptionEn"`
			} `json:"conNumberTypes"`
			ConCode        string `json:"conCode"`
			ConAreaCode    string `json:"conAreaCode"`
			ConPhoneNumber string `json:"conPhoneNumber"`
			ConExtension   string `json:"conExtension"`
		} `json:"contacts"`
		Employers []struct {
			EmpEmployerNameDescAr string `json:"empEmployerNameDescAr"`
			EmpEmployerNameDescEn string `json:"empEmployerNameDescEn"`
			EmpOccupationDescAr   string `json:"empOccupationDescAr"`
			EmpOccupationDescEn   string `json:"empOccupationDescEn"`
			EmpDateOfEmployment   string `json:"empDateOfEmployment"`
			EmpDateLoaded         string `json:"empDateLoaded"`
			EmpIncome             int    `json:"empIncome"`
			EmpTotalIncome        int    `json:"empTotalIncome"`
			EmpAddress            struct {
				AdrsDateLoaded   string `json:"adrsDateLoaded"`
				AdrsAddressTypes struct {
					AddressID       int    `json:"addressID"`
					AddressTypeCode string `json:"addressTypeCode"`
					AddressNameAR   string `json:"addressNameAR"`
					AddressNameEN   string `json:"addressNameEN"`
				} `json:"adrsAddressTypes"`
				AdrsAddressLineFirstDescAr  string `json:"adrsAddressLineFirstDescAr"`
				AdrsAddressLineFirstDescEn  string `json:"adrsAddressLineFirstDescEn"`
				AdrsAddressLineSecondDescAr string `json:"adrsAddressLineSecondDescAr"`
				AdrsAddressLineSecondDescEn string `json:"adrsAddressLineSecondDescEn"`
				AdrsPOBox                   string `json:"adrsPOBox"`
				AdrsPostalCode              string `json:"adrsPostalCode"`
				AdrsCityDescAr              string `json:"adrsCityDescAr"`
				AdrsCityDescEn              string `json:"adrsCityDescEn"`
				NationalAddress             struct {
					BuildingNumber   string `json:"buildingNumber"`
					StreetAr         string `json:"streetAr"`
					StreetEn         string `json:"streetEn"`
					DistrictAr       string `json:"districtAr"`
					DistrictEn       string `json:"districtEn"`
					AdditionalNumber string `json:"additionalNumber"`
					UnitNumber       string `json:"unitNumber"`
				} `json:"nationalAddress"`
			} `json:"empAddress"`
			EmpStatusType struct {
				EmployerStatusTypeCode   string `json:"employerStatusTypeCode"`
				EmployerStatusTypeDescAr string `json:"employerStatusTypeDescAr"`
				EmployerStatusTypeDescEn string `json:"employerStatusTypeDescEn"`
			} `json:"empStatusType"`
		} `json:"employers"`
		SummaryInfo struct {
			SummActiveCreditInstruments     int    `json:"summActiveCreditInstruments"`
			SummDefaults                    int    `json:"summDefaults"`
			SummEarliestIssueDate           string `json:"summEarliestIssueDate"`
			SummTotalLimits                 int    `json:"summTotalLimits"`
			SummTotalGuaranteedLimits       int    `json:"summTotalGuaranteedLimits"`
			SummTotalLiablilites            int    `json:"summTotalLiablilites"`
			SummTotalGuaranteedLiablilites  int    `json:"summTotalGuaranteedLiablilites"`
			SummTotalDefaults               int    `json:"summTotalDefaults"`
			SummCurrentDelinquentBalance    int    `json:"summCurrentDelinquentBalance"`
			SummPreviousEnquires            int    `json:"summPreviousEnquires"`
			SummPreviousEnquiresThisMonth   int    `json:"summPreviousEnquiresThisMonth"`
			SummGuaranteedCreditInstruments int    `json:"summGuaranteedCreditInstruments"`
		} `json:"summaryInfo"`
		DisclerText struct {
			DiscTextDescAr string `json:"discTextDescAr"`
			DiscTextDescEn string `json:"discTextDescEn"`
		} `json:"disclerText"`
		Score []struct {
			Score       int `json:"score"`
			ReasonCodes []struct {
				ScoreReasonCodeName   string `json:"scoreReasonCodeName"`
				ScoreReasonCodeDescAr string `json:"scoreReasonCodeDescAr"`
				ScoreReasonCodeDescEn string `json:"scoreReasonCodeDescEn"`
			} `json:"reasonCodes"`
			ScoreCard struct {
				ScoreCardCode   string `json:"scoreCardCode"`
				ScoreCardDescAr string `json:"scoreCardDescAr"`
				ScoreCardDescEn string `json:"scoreCardDescEn"`
			} `json:"scoreCard"`
			MinimumScore int    `json:"minimumScore"`
			MaximumScore int    `json:"maximumScore"`
			ScoreIndex   int    `json:"scoreIndex"`
			Error        string `json:"error"`
		} `json:"score"`
		AffordabilityResponseModel struct {
			TotalIncome              int  `json:"totalIncome"`
			TotalCreditCommitment    int  `json:"totalCreditCommitment"`
			TotalNonCreditCommitment int  `json:"totalNonCreditCommitment"`
			DisposableIncome         int  `json:"disposableIncome"`
			IsEligible               bool `json:"isEligible"`
			ProductAffordability     struct {
				SalariedNonMortgage struct {
					ApplicableDBR           int `json:"applicableDBR"`
					MonthlyInstallmentLimit int `json:"monthlyInstallmentLimit"`
				} `json:"salariedNonMortgage"`
				NonSalariedNonMortgage struct {
					ApplicableDBR           int `json:"applicableDBR"`
					MonthlyInstallmentLimit int `json:"monthlyInstallmentLimit"`
				} `json:"nonSalariedNonMortgage"`
				NonRedfMortgage struct {
					ApplicableDBR           int `json:"applicableDBR"`
					MonthlyInstallmentLimit int `json:"monthlyInstallmentLimit"`
				} `json:"nonRedfMortgage"`
				RedfMortgage struct {
					ApplicableDBR           int `json:"applicableDBR"`
					MonthlyInstallmentLimit int `json:"monthlyInstallmentLimit"`
				} `json:"redfMortgage"`
				GlobalDBR           int `json:"globalDBR"`
				DbrForSalariedLoans int `json:"dbrForSalariedLoans"`
				ApplicableDBR       int `json:"applicableDBR"`
			} `json:"productAffordability"`
			TotalValueUsedIncaluculation int `json:"totalValueUsedIncaluculation"`
			TotalValueDeclaredByCustomer int `json:"totalValueDeclaredByCustomer"`
			TotalOutputValue             int `json:"totalOutputValue"`
			ExpenseResponseModel         []struct {
				NameEn                  string `json:"nameEn"`
				NameAr                  string `json:"nameAr"`
				ValueUsedIncaluculation int    `json:"valueUsedIncaluculation"`
				ValueDeclaredByCustomer int    `json:"valueDeclaredByCustomer"`
				OutputValue             int    `json:"outputValue"`
				IsVerified              bool   `json:"isVerified"`
			} `json:"expenseResponseModel"`
			UserInput struct {
				BasicIncome      int  `json:"basicIncome"`
				AdditionalIncome int  `json:"additionalIncome"`
				IsRetired        bool `json:"isRetired"`
				Breadwinner      bool `json:"breadwinner"`
				CityOfResidence  struct {
					TextEn string `json:"textEn"`
					TextAr string `json:"textAr"`
				} `json:"cityOfResidence"`
				Nationality struct {
					TextEn string `json:"textEn"`
					TextAr string `json:"textAr"`
				} `json:"nationality"`
				MaritalStatus struct {
					TextEn string `json:"textEn"`
					TextAr string `json:"textAr"`
				} `json:"maritalStatus"`
				NumberOfDependents                int `json:"numberOfDependents"`
				NumberOfDependentsInPrivateSchool int `json:"numberOfDependentsInPrivateSchool"`
				NumberOfDependentsInPublicSchool  int `json:"numberOfDependentsInPublicSchool"`
				NumberOfHouseholdhelp             int `json:"numberOfHouseholdhelp"`
				HomeOwnership                     struct {
					TextEn string `json:"textEn"`
					TextAr string `json:"textAr"`
				} `json:"homeOwnership"`
				ResidentialType struct {
					TextEn string `json:"textEn"`
					TextAr string `json:"textAr"`
				} `json:"residentialType"`
				IsRedfCustomer bool `json:"isRedfCustomer"`
				Expenses       []struct {
					NameEn                  string `json:"nameEn"`
					NameAr                  string `json:"nameAr"`
					ValueUsedIncaluculation int    `json:"valueUsedIncaluculation"`
					ValueDeclaredByCustomer int    `json:"valueDeclaredByCustomer"`
					OutputValue             int    `json:"outputValue"`
					IsVerified              bool   `json:"isVerified"`
				} `json:"expenses"`
			} `json:"userInput"`
		} `json:"affordabilityResponseModel"`
	} `json:"data"`
}
