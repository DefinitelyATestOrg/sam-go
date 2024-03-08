// File generated from our OpenAPI spec by Stainless.

package samgo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/internal/apiquery"
	"github.com/DefinitelyATestOrg/sam-go/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/option"
)

// CustomerAccountService contains methods and other services that help with
// interacting with the sam API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCustomerAccountService] method
// instead.
type CustomerAccountService struct {
	Options []option.RequestOption
}

// NewCustomerAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerAccountService(opts ...option.RequestOption) (r *CustomerAccountService) {
	r = &CustomerAccountService{}
	r.Options = opts
	return
}

// Allows retrieval of a single account, fetched by its id.
func (r *CustomerAccountService) Get(ctx context.Context, customerID string, accountID string, query CustomerAccountGetParams, opts ...option.RequestOption) (res *CustomerAccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/customers/%s/accounts/%s", customerID, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Allows retrieval of accounts associated with the current authorised user. It's
// possible to filter the accounts list by status or by cashAccountType.
func (r *CustomerAccountService) List(ctx context.Context, customerID string, params CustomerAccountListParams, opts ...option.RequestOption) (res *CustomerAccountListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/customers/%s/accounts", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Allows closure of a single account.
func (r *CustomerAccountService) Close(ctx context.Context, customerID string, accountID string, body CustomerAccountCloseParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("v1/customers/%s/accounts/%s/close", customerID, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomerAccountGetResponse struct {
	Account CustomerAccountGetResponseAccount `json:"account"`
	JSON    customerAccountGetResponseJSON    `json:"-"`
}

// customerAccountGetResponseJSON contains the JSON metadata for the struct
// [CustomerAccountGetResponse]
type customerAccountGetResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerAccountGetResponseAccount struct {
	// ISO 4217 Alpha 3 currency code.
	Currency   string `json:"currency,required"`
	ResourceID string `json:"resourceId,required"`
	// Account status. The value is one of the following:
	//
	// - "enabled": account is available
	// - "deleted": account is terminated
	// - "blocked": account is blocked e.g. for legal reasons
	Status CustomerAccountGetResponseAccountStatus `json:"status,required"`
	// A list of balances regarding this account, e.g. the current balance, the last
	// booked balance. The list might be restricted to the current balance.
	Balances []CustomerAccountGetResponseAccountBalance `json:"balances"`
	// Basic Bank Account Number (BBAN) Identifier.
	Bban string `json:"bban"`
	// BICFI
	Bic string `json:"bic"`
	// ExternalCashAccountType1Code from ISO 20022.
	CashAccountType CustomerAccountGetResponseAccountCashAccountType `json:"cashAccountType"`
	// Additional information on the account.
	Details string `json:"details"`
	// Name of the account as defined by the user within online channels.
	DisplayName string `json:"displayName"`
	// IBAN of an account.
	Iban string `json:"iban"`
	// Mobile phone number.
	Msisdn string `json:"msisdn"`
	// Name of the account.
	Name string `json:"name"`
	// Name of the legal account owner. If there is more than one owner, then e.g. two
	// names might be noted here. For a corporate account, the corporate name is used
	// for this attribute.
	OwnerName string `json:"ownerName"`
	// List of owner names.
	OwnerNames []CustomerAccountGetResponseAccountOwnerName `json:"ownerNames"`
	// Product name of the bank for this account, proprietary definition.
	Product string `json:"product"`
	// Name of the user. In case of a corporate account, this might be the person
	// acting on behalf of the corporate.
	PsuName string `json:"psuName"`
	// Specifies the usage of the account:
	//
	// - PRIV: private personal account
	// - ORGA: professional account
	Usage CustomerAccountGetResponseAccountUsage `json:"usage"`
	JSON  customerAccountGetResponseAccountJSON  `json:"-"`
}

// customerAccountGetResponseAccountJSON contains the JSON metadata for the struct
// [CustomerAccountGetResponseAccount]
type customerAccountGetResponseAccountJSON struct {
	Currency        apijson.Field
	ResourceID      apijson.Field
	Status          apijson.Field
	Balances        apijson.Field
	Bban            apijson.Field
	Bic             apijson.Field
	CashAccountType apijson.Field
	Details         apijson.Field
	DisplayName     apijson.Field
	Iban            apijson.Field
	Msisdn          apijson.Field
	Name            apijson.Field
	OwnerName       apijson.Field
	OwnerNames      apijson.Field
	Product         apijson.Field
	PsuName         apijson.Field
	Usage           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomerAccountGetResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountGetResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Account status. The value is one of the following:
//
// - "enabled": account is available
// - "deleted": account is terminated
// - "blocked": account is blocked e.g. for legal reasons
type CustomerAccountGetResponseAccountStatus string

const (
	CustomerAccountGetResponseAccountStatusEnabled CustomerAccountGetResponseAccountStatus = "enabled"
	CustomerAccountGetResponseAccountStatusDeleted CustomerAccountGetResponseAccountStatus = "deleted"
	CustomerAccountGetResponseAccountStatusBlocked CustomerAccountGetResponseAccountStatus = "blocked"
)

// A single balance element.
type CustomerAccountGetResponseAccountBalance struct {
	BalanceAmount CustomerAccountGetResponseAccountBalancesBalanceAmount `json:"balanceAmount,required"`
	// The following balance types are defined:
	//
	//   - "closingBooked": Balance of the account at the end of the pre-agreed account
	//     reporting period. It is the sum of the opening booked balance at the beginning
	//     of the period and all entries booked to the account during the pre-agreed
	//     account reporting period.
	//   - "openingBooked": Book balance of the account at the beginning of the account
	//     reporting period. It always equals the closing book balance from the previous
	//     report.
	//   - "interimAvailable": Available balance calculated in the course of the account
	//     ?servicer?s business day, at the time specified, and subject to further
	//     changes during the business day. The interim balance is calculated on the
	//     basis of booked credit and debit items during the calculation time/period
	//     specified.
	//   - "interimBooked": Balance calculated in the course of the account servicer's
	//     business day, at the time specified, and subject to further changes during the
	//     business day. The interim balance is calculated on the basis of booked credit
	//     and debit items during the calculation time/period specified.
	BalanceType CustomerAccountGetResponseAccountBalancesBalanceType `json:"balanceType,required"`
	// A flag indicating if the credit limit of the corresponding account is included
	// in the calculation of the balance, where applicable.
	CreditLimitIncluded bool `json:"creditLimitIncluded"`
	// Indicates the date and time of the balance.
	LastRetrievedDateTime time.Time                                    `json:"lastRetrievedDateTime" format:"date-time"`
	JSON                  customerAccountGetResponseAccountBalanceJSON `json:"-"`
}

// customerAccountGetResponseAccountBalanceJSON contains the JSON metadata for the
// struct [CustomerAccountGetResponseAccountBalance]
type customerAccountGetResponseAccountBalanceJSON struct {
	BalanceAmount         apijson.Field
	BalanceType           apijson.Field
	CreditLimitIncluded   apijson.Field
	LastRetrievedDateTime apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CustomerAccountGetResponseAccountBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountGetResponseAccountBalanceJSON) RawJSON() string {
	return r.raw
}

type CustomerAccountGetResponseAccountBalancesBalanceAmount struct {
	// The amount given with fractional digits, where fractions must be compliant to
	// the currency definition. Up to 14 significant figures. Negative amounts are
	// signed by minus. The decimal separator is a dot.
	//
	// **Example:** Valid representations for EUR with up to two decimals are:
	//
	// - 1056
	// - 5768.2
	// - -1.50
	// - 5877.78
	Amount string `json:"amount,required"`
	// ISO 4217 Alpha 3 currency code.
	Currency string                                                     `json:"currency,required"`
	JSON     customerAccountGetResponseAccountBalancesBalanceAmountJSON `json:"-"`
}

// customerAccountGetResponseAccountBalancesBalanceAmountJSON contains the JSON
// metadata for the struct [CustomerAccountGetResponseAccountBalancesBalanceAmount]
type customerAccountGetResponseAccountBalancesBalanceAmountJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAccountGetResponseAccountBalancesBalanceAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountGetResponseAccountBalancesBalanceAmountJSON) RawJSON() string {
	return r.raw
}

// The following balance types are defined:
//
//   - "closingBooked": Balance of the account at the end of the pre-agreed account
//     reporting period. It is the sum of the opening booked balance at the beginning
//     of the period and all entries booked to the account during the pre-agreed
//     account reporting period.
//   - "openingBooked": Book balance of the account at the beginning of the account
//     reporting period. It always equals the closing book balance from the previous
//     report.
//   - "interimAvailable": Available balance calculated in the course of the account
//     ?servicer?s business day, at the time specified, and subject to further
//     changes during the business day. The interim balance is calculated on the
//     basis of booked credit and debit items during the calculation time/period
//     specified.
//   - "interimBooked": Balance calculated in the course of the account servicer's
//     business day, at the time specified, and subject to further changes during the
//     business day. The interim balance is calculated on the basis of booked credit
//     and debit items during the calculation time/period specified.
type CustomerAccountGetResponseAccountBalancesBalanceType string

const (
	CustomerAccountGetResponseAccountBalancesBalanceTypeClosingBooked    CustomerAccountGetResponseAccountBalancesBalanceType = "closingBooked"
	CustomerAccountGetResponseAccountBalancesBalanceTypeOpeningBooked    CustomerAccountGetResponseAccountBalancesBalanceType = "openingBooked"
	CustomerAccountGetResponseAccountBalancesBalanceTypeInterimAvailable CustomerAccountGetResponseAccountBalancesBalanceType = "interimAvailable"
	CustomerAccountGetResponseAccountBalancesBalanceTypeInterimBooked    CustomerAccountGetResponseAccountBalancesBalanceType = "interimBooked"
)

// ExternalCashAccountType1Code from ISO 20022.
type CustomerAccountGetResponseAccountCashAccountType string

const (
	CustomerAccountGetResponseAccountCashAccountTypeCacc CustomerAccountGetResponseAccountCashAccountType = "CACC"
	CustomerAccountGetResponseAccountCashAccountTypeCard CustomerAccountGetResponseAccountCashAccountType = "CARD"
	CustomerAccountGetResponseAccountCashAccountTypeCash CustomerAccountGetResponseAccountCashAccountType = "CASH"
	CustomerAccountGetResponseAccountCashAccountTypeChar CustomerAccountGetResponseAccountCashAccountType = "CHAR"
	CustomerAccountGetResponseAccountCashAccountTypeCish CustomerAccountGetResponseAccountCashAccountType = "CISH"
	CustomerAccountGetResponseAccountCashAccountTypeComm CustomerAccountGetResponseAccountCashAccountType = "COMM"
	CustomerAccountGetResponseAccountCashAccountTypeCpac CustomerAccountGetResponseAccountCashAccountType = "CPAC"
	CustomerAccountGetResponseAccountCashAccountTypeLlsv CustomerAccountGetResponseAccountCashAccountType = "LLSV"
	CustomerAccountGetResponseAccountCashAccountTypeLoan CustomerAccountGetResponseAccountCashAccountType = "LOAN"
	CustomerAccountGetResponseAccountCashAccountTypeMgld CustomerAccountGetResponseAccountCashAccountType = "MGLD"
	CustomerAccountGetResponseAccountCashAccountTypeMoma CustomerAccountGetResponseAccountCashAccountType = "MOMA"
	CustomerAccountGetResponseAccountCashAccountTypeNrex CustomerAccountGetResponseAccountCashAccountType = "NREX"
	CustomerAccountGetResponseAccountCashAccountTypeOdft CustomerAccountGetResponseAccountCashAccountType = "ODFT"
	CustomerAccountGetResponseAccountCashAccountTypeOndp CustomerAccountGetResponseAccountCashAccountType = "ONDP"
	CustomerAccountGetResponseAccountCashAccountTypeOthr CustomerAccountGetResponseAccountCashAccountType = "OTHR"
	CustomerAccountGetResponseAccountCashAccountTypeSacc CustomerAccountGetResponseAccountCashAccountType = "SACC"
	CustomerAccountGetResponseAccountCashAccountTypeSlry CustomerAccountGetResponseAccountCashAccountType = "SLRY"
	CustomerAccountGetResponseAccountCashAccountTypeSvgs CustomerAccountGetResponseAccountCashAccountType = "SVGS"
	CustomerAccountGetResponseAccountCashAccountTypeTaxe CustomerAccountGetResponseAccountCashAccountType = "TAXE"
	CustomerAccountGetResponseAccountCashAccountTypeTran CustomerAccountGetResponseAccountCashAccountType = "TRAN"
	CustomerAccountGetResponseAccountCashAccountTypeTras CustomerAccountGetResponseAccountCashAccountType = "TRAS"
	CustomerAccountGetResponseAccountCashAccountTypeVacc CustomerAccountGetResponseAccountCashAccountType = "VACC"
	CustomerAccountGetResponseAccountCashAccountTypeNfca CustomerAccountGetResponseAccountCashAccountType = "NFCA"
)

// Identifies a Person that acts as an account owner
type CustomerAccountGetResponseAccountOwnerName struct {
	// Account owner name
	Name string `json:"name,required"`
	// The following proprietary codes are used:
	//
	// - "owner",
	// - "legalRepresentative",
	// - "authorisedUser"
	Role string                                         `json:"role"`
	JSON customerAccountGetResponseAccountOwnerNameJSON `json:"-"`
}

// customerAccountGetResponseAccountOwnerNameJSON contains the JSON metadata for
// the struct [CustomerAccountGetResponseAccountOwnerName]
type customerAccountGetResponseAccountOwnerNameJSON struct {
	Name        apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAccountGetResponseAccountOwnerName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountGetResponseAccountOwnerNameJSON) RawJSON() string {
	return r.raw
}

// Specifies the usage of the account:
//
// - PRIV: private personal account
// - ORGA: professional account
type CustomerAccountGetResponseAccountUsage string

const (
	CustomerAccountGetResponseAccountUsagePriv CustomerAccountGetResponseAccountUsage = "PRIV"
	CustomerAccountGetResponseAccountUsageOrga CustomerAccountGetResponseAccountUsage = "ORGA"
)

type CustomerAccountListResponse struct {
	Accounts []CustomerAccountListResponseAccount `json:"accounts"`
	JSON     customerAccountListResponseJSON      `json:"-"`
}

// customerAccountListResponseJSON contains the JSON metadata for the struct
// [CustomerAccountListResponse]
type customerAccountListResponseJSON struct {
	Accounts    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerAccountListResponseAccount struct {
	// ISO 4217 Alpha 3 currency code.
	Currency   string `json:"currency,required"`
	ResourceID string `json:"resourceId,required"`
	// Account status. The value is one of the following:
	//
	// - "enabled": account is available
	// - "deleted": account is terminated
	// - "blocked": account is blocked e.g. for legal reasons
	Status CustomerAccountListResponseAccountsStatus `json:"status,required"`
	// A list of balances regarding this account, e.g. the current balance, the last
	// booked balance. The list might be restricted to the current balance.
	Balances []CustomerAccountListResponseAccountsBalance `json:"balances"`
	// Basic Bank Account Number (BBAN) Identifier.
	Bban string `json:"bban"`
	// BICFI
	Bic string `json:"bic"`
	// ExternalCashAccountType1Code from ISO 20022.
	CashAccountType CustomerAccountListResponseAccountsCashAccountType `json:"cashAccountType"`
	// Additional information on the account.
	Details string `json:"details"`
	// Name of the account as defined by the user within online channels.
	DisplayName string `json:"displayName"`
	// IBAN of an account.
	Iban string `json:"iban"`
	// Mobile phone number.
	Msisdn string `json:"msisdn"`
	// Name of the account.
	Name string `json:"name"`
	// Name of the legal account owner. If there is more than one owner, then e.g. two
	// names might be noted here. For a corporate account, the corporate name is used
	// for this attribute.
	OwnerName string `json:"ownerName"`
	// List of owner names.
	OwnerNames []CustomerAccountListResponseAccountsOwnerName `json:"ownerNames"`
	// Product name of the bank for this account, proprietary definition.
	Product string `json:"product"`
	// Name of the user. In case of a corporate account, this might be the person
	// acting on behalf of the corporate.
	PsuName string `json:"psuName"`
	// Specifies the usage of the account:
	//
	// - PRIV: private personal account
	// - ORGA: professional account
	Usage CustomerAccountListResponseAccountsUsage `json:"usage"`
	JSON  customerAccountListResponseAccountJSON   `json:"-"`
}

// customerAccountListResponseAccountJSON contains the JSON metadata for the struct
// [CustomerAccountListResponseAccount]
type customerAccountListResponseAccountJSON struct {
	Currency        apijson.Field
	ResourceID      apijson.Field
	Status          apijson.Field
	Balances        apijson.Field
	Bban            apijson.Field
	Bic             apijson.Field
	CashAccountType apijson.Field
	Details         apijson.Field
	DisplayName     apijson.Field
	Iban            apijson.Field
	Msisdn          apijson.Field
	Name            apijson.Field
	OwnerName       apijson.Field
	OwnerNames      apijson.Field
	Product         apijson.Field
	PsuName         apijson.Field
	Usage           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomerAccountListResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountListResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Account status. The value is one of the following:
//
// - "enabled": account is available
// - "deleted": account is terminated
// - "blocked": account is blocked e.g. for legal reasons
type CustomerAccountListResponseAccountsStatus string

const (
	CustomerAccountListResponseAccountsStatusEnabled CustomerAccountListResponseAccountsStatus = "enabled"
	CustomerAccountListResponseAccountsStatusDeleted CustomerAccountListResponseAccountsStatus = "deleted"
	CustomerAccountListResponseAccountsStatusBlocked CustomerAccountListResponseAccountsStatus = "blocked"
)

// A single balance element.
type CustomerAccountListResponseAccountsBalance struct {
	BalanceAmount CustomerAccountListResponseAccountsBalancesBalanceAmount `json:"balanceAmount,required"`
	// The following balance types are defined:
	//
	//   - "closingBooked": Balance of the account at the end of the pre-agreed account
	//     reporting period. It is the sum of the opening booked balance at the beginning
	//     of the period and all entries booked to the account during the pre-agreed
	//     account reporting period.
	//   - "openingBooked": Book balance of the account at the beginning of the account
	//     reporting period. It always equals the closing book balance from the previous
	//     report.
	//   - "interimAvailable": Available balance calculated in the course of the account
	//     ?servicer?s business day, at the time specified, and subject to further
	//     changes during the business day. The interim balance is calculated on the
	//     basis of booked credit and debit items during the calculation time/period
	//     specified.
	//   - "interimBooked": Balance calculated in the course of the account servicer's
	//     business day, at the time specified, and subject to further changes during the
	//     business day. The interim balance is calculated on the basis of booked credit
	//     and debit items during the calculation time/period specified.
	BalanceType CustomerAccountListResponseAccountsBalancesBalanceType `json:"balanceType,required"`
	// A flag indicating if the credit limit of the corresponding account is included
	// in the calculation of the balance, where applicable.
	CreditLimitIncluded bool `json:"creditLimitIncluded"`
	// Indicates the date and time of the balance.
	LastRetrievedDateTime time.Time                                      `json:"lastRetrievedDateTime" format:"date-time"`
	JSON                  customerAccountListResponseAccountsBalanceJSON `json:"-"`
}

// customerAccountListResponseAccountsBalanceJSON contains the JSON metadata for
// the struct [CustomerAccountListResponseAccountsBalance]
type customerAccountListResponseAccountsBalanceJSON struct {
	BalanceAmount         apijson.Field
	BalanceType           apijson.Field
	CreditLimitIncluded   apijson.Field
	LastRetrievedDateTime apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CustomerAccountListResponseAccountsBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountListResponseAccountsBalanceJSON) RawJSON() string {
	return r.raw
}

type CustomerAccountListResponseAccountsBalancesBalanceAmount struct {
	// The amount given with fractional digits, where fractions must be compliant to
	// the currency definition. Up to 14 significant figures. Negative amounts are
	// signed by minus. The decimal separator is a dot.
	//
	// **Example:** Valid representations for EUR with up to two decimals are:
	//
	// - 1056
	// - 5768.2
	// - -1.50
	// - 5877.78
	Amount string `json:"amount,required"`
	// ISO 4217 Alpha 3 currency code.
	Currency string                                                       `json:"currency,required"`
	JSON     customerAccountListResponseAccountsBalancesBalanceAmountJSON `json:"-"`
}

// customerAccountListResponseAccountsBalancesBalanceAmountJSON contains the JSON
// metadata for the struct
// [CustomerAccountListResponseAccountsBalancesBalanceAmount]
type customerAccountListResponseAccountsBalancesBalanceAmountJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAccountListResponseAccountsBalancesBalanceAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountListResponseAccountsBalancesBalanceAmountJSON) RawJSON() string {
	return r.raw
}

// The following balance types are defined:
//
//   - "closingBooked": Balance of the account at the end of the pre-agreed account
//     reporting period. It is the sum of the opening booked balance at the beginning
//     of the period and all entries booked to the account during the pre-agreed
//     account reporting period.
//   - "openingBooked": Book balance of the account at the beginning of the account
//     reporting period. It always equals the closing book balance from the previous
//     report.
//   - "interimAvailable": Available balance calculated in the course of the account
//     ?servicer?s business day, at the time specified, and subject to further
//     changes during the business day. The interim balance is calculated on the
//     basis of booked credit and debit items during the calculation time/period
//     specified.
//   - "interimBooked": Balance calculated in the course of the account servicer's
//     business day, at the time specified, and subject to further changes during the
//     business day. The interim balance is calculated on the basis of booked credit
//     and debit items during the calculation time/period specified.
type CustomerAccountListResponseAccountsBalancesBalanceType string

const (
	CustomerAccountListResponseAccountsBalancesBalanceTypeClosingBooked    CustomerAccountListResponseAccountsBalancesBalanceType = "closingBooked"
	CustomerAccountListResponseAccountsBalancesBalanceTypeOpeningBooked    CustomerAccountListResponseAccountsBalancesBalanceType = "openingBooked"
	CustomerAccountListResponseAccountsBalancesBalanceTypeInterimAvailable CustomerAccountListResponseAccountsBalancesBalanceType = "interimAvailable"
	CustomerAccountListResponseAccountsBalancesBalanceTypeInterimBooked    CustomerAccountListResponseAccountsBalancesBalanceType = "interimBooked"
)

// ExternalCashAccountType1Code from ISO 20022.
type CustomerAccountListResponseAccountsCashAccountType string

const (
	CustomerAccountListResponseAccountsCashAccountTypeCacc CustomerAccountListResponseAccountsCashAccountType = "CACC"
	CustomerAccountListResponseAccountsCashAccountTypeCard CustomerAccountListResponseAccountsCashAccountType = "CARD"
	CustomerAccountListResponseAccountsCashAccountTypeCash CustomerAccountListResponseAccountsCashAccountType = "CASH"
	CustomerAccountListResponseAccountsCashAccountTypeChar CustomerAccountListResponseAccountsCashAccountType = "CHAR"
	CustomerAccountListResponseAccountsCashAccountTypeCish CustomerAccountListResponseAccountsCashAccountType = "CISH"
	CustomerAccountListResponseAccountsCashAccountTypeComm CustomerAccountListResponseAccountsCashAccountType = "COMM"
	CustomerAccountListResponseAccountsCashAccountTypeCpac CustomerAccountListResponseAccountsCashAccountType = "CPAC"
	CustomerAccountListResponseAccountsCashAccountTypeLlsv CustomerAccountListResponseAccountsCashAccountType = "LLSV"
	CustomerAccountListResponseAccountsCashAccountTypeLoan CustomerAccountListResponseAccountsCashAccountType = "LOAN"
	CustomerAccountListResponseAccountsCashAccountTypeMgld CustomerAccountListResponseAccountsCashAccountType = "MGLD"
	CustomerAccountListResponseAccountsCashAccountTypeMoma CustomerAccountListResponseAccountsCashAccountType = "MOMA"
	CustomerAccountListResponseAccountsCashAccountTypeNrex CustomerAccountListResponseAccountsCashAccountType = "NREX"
	CustomerAccountListResponseAccountsCashAccountTypeOdft CustomerAccountListResponseAccountsCashAccountType = "ODFT"
	CustomerAccountListResponseAccountsCashAccountTypeOndp CustomerAccountListResponseAccountsCashAccountType = "ONDP"
	CustomerAccountListResponseAccountsCashAccountTypeOthr CustomerAccountListResponseAccountsCashAccountType = "OTHR"
	CustomerAccountListResponseAccountsCashAccountTypeSacc CustomerAccountListResponseAccountsCashAccountType = "SACC"
	CustomerAccountListResponseAccountsCashAccountTypeSlry CustomerAccountListResponseAccountsCashAccountType = "SLRY"
	CustomerAccountListResponseAccountsCashAccountTypeSvgs CustomerAccountListResponseAccountsCashAccountType = "SVGS"
	CustomerAccountListResponseAccountsCashAccountTypeTaxe CustomerAccountListResponseAccountsCashAccountType = "TAXE"
	CustomerAccountListResponseAccountsCashAccountTypeTran CustomerAccountListResponseAccountsCashAccountType = "TRAN"
	CustomerAccountListResponseAccountsCashAccountTypeTras CustomerAccountListResponseAccountsCashAccountType = "TRAS"
	CustomerAccountListResponseAccountsCashAccountTypeVacc CustomerAccountListResponseAccountsCashAccountType = "VACC"
	CustomerAccountListResponseAccountsCashAccountTypeNfca CustomerAccountListResponseAccountsCashAccountType = "NFCA"
)

// Identifies a Person that acts as an account owner
type CustomerAccountListResponseAccountsOwnerName struct {
	// Account owner name
	Name string `json:"name,required"`
	// The following proprietary codes are used:
	//
	// - "owner",
	// - "legalRepresentative",
	// - "authorisedUser"
	Role string                                           `json:"role"`
	JSON customerAccountListResponseAccountsOwnerNameJSON `json:"-"`
}

// customerAccountListResponseAccountsOwnerNameJSON contains the JSON metadata for
// the struct [CustomerAccountListResponseAccountsOwnerName]
type customerAccountListResponseAccountsOwnerNameJSON struct {
	Name        apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAccountListResponseAccountsOwnerName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountListResponseAccountsOwnerNameJSON) RawJSON() string {
	return r.raw
}

// Specifies the usage of the account:
//
// - PRIV: private personal account
// - ORGA: professional account
type CustomerAccountListResponseAccountsUsage string

const (
	CustomerAccountListResponseAccountsUsagePriv CustomerAccountListResponseAccountsUsage = "PRIV"
	CustomerAccountListResponseAccountsUsageOrga CustomerAccountListResponseAccountsUsage = "ORGA"
)

type CustomerAccountGetParams struct {
	UserID param.Field[string] `header:"userId,required"`
}

type CustomerAccountListParams struct {
	UserID param.Field[string] `header:"userId,required"`
	// A filter on the list based on the cashAccountType field. The value must be a 4
	// characters string as defined in ISO 20022
	CashAccountType param.Field[[]CustomerAccountListParamsCashAccountType] `query:"cashAccountType"`
	// A filter on the list based on the account status field. Available values:
	//
	// - enabled - account is available
	// - deleted - account is terminated
	// - blocked - account is blocked e.g. for legal reasons
	Status param.Field[[]CustomerAccountListParamsStatus] `query:"status"`
}

// URLQuery serializes [CustomerAccountListParams]'s query parameters as
// `url.Values`.
func (r CustomerAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ExternalCashAccountType1Code from ISO 20022.
type CustomerAccountListParamsCashAccountType string

const (
	CustomerAccountListParamsCashAccountTypeCacc CustomerAccountListParamsCashAccountType = "CACC"
	CustomerAccountListParamsCashAccountTypeCard CustomerAccountListParamsCashAccountType = "CARD"
	CustomerAccountListParamsCashAccountTypeCash CustomerAccountListParamsCashAccountType = "CASH"
	CustomerAccountListParamsCashAccountTypeChar CustomerAccountListParamsCashAccountType = "CHAR"
	CustomerAccountListParamsCashAccountTypeCish CustomerAccountListParamsCashAccountType = "CISH"
	CustomerAccountListParamsCashAccountTypeComm CustomerAccountListParamsCashAccountType = "COMM"
	CustomerAccountListParamsCashAccountTypeCpac CustomerAccountListParamsCashAccountType = "CPAC"
	CustomerAccountListParamsCashAccountTypeLlsv CustomerAccountListParamsCashAccountType = "LLSV"
	CustomerAccountListParamsCashAccountTypeLoan CustomerAccountListParamsCashAccountType = "LOAN"
	CustomerAccountListParamsCashAccountTypeMgld CustomerAccountListParamsCashAccountType = "MGLD"
	CustomerAccountListParamsCashAccountTypeMoma CustomerAccountListParamsCashAccountType = "MOMA"
	CustomerAccountListParamsCashAccountTypeNrex CustomerAccountListParamsCashAccountType = "NREX"
	CustomerAccountListParamsCashAccountTypeOdft CustomerAccountListParamsCashAccountType = "ODFT"
	CustomerAccountListParamsCashAccountTypeOndp CustomerAccountListParamsCashAccountType = "ONDP"
	CustomerAccountListParamsCashAccountTypeOthr CustomerAccountListParamsCashAccountType = "OTHR"
	CustomerAccountListParamsCashAccountTypeSacc CustomerAccountListParamsCashAccountType = "SACC"
	CustomerAccountListParamsCashAccountTypeSlry CustomerAccountListParamsCashAccountType = "SLRY"
	CustomerAccountListParamsCashAccountTypeSvgs CustomerAccountListParamsCashAccountType = "SVGS"
	CustomerAccountListParamsCashAccountTypeTaxe CustomerAccountListParamsCashAccountType = "TAXE"
	CustomerAccountListParamsCashAccountTypeTran CustomerAccountListParamsCashAccountType = "TRAN"
	CustomerAccountListParamsCashAccountTypeTras CustomerAccountListParamsCashAccountType = "TRAS"
	CustomerAccountListParamsCashAccountTypeVacc CustomerAccountListParamsCashAccountType = "VACC"
	CustomerAccountListParamsCashAccountTypeNfca CustomerAccountListParamsCashAccountType = "NFCA"
)

// Account status. The value is one of the following:
//
// - "enabled": account is available
// - "deleted": account is terminated
// - "blocked": account is blocked e.g. for legal reasons
type CustomerAccountListParamsStatus string

const (
	CustomerAccountListParamsStatusEnabled CustomerAccountListParamsStatus = "enabled"
	CustomerAccountListParamsStatusDeleted CustomerAccountListParamsStatus = "deleted"
	CustomerAccountListParamsStatusBlocked CustomerAccountListParamsStatus = "blocked"
)

type CustomerAccountCloseParams struct {
	IdempotencyKey param.Field[string] `header:"idempotencyKey"`
}
