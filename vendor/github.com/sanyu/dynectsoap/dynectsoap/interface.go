package dynectsoap

type Dynect interface {

	// Error can be either of the following types:
	//
	//   - fault

	GetAllRecords(request *GetAllRecordsRequestType) (*GetAllRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetJob(request *GetJobRequestType) (*GetJobResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetJobRetry(request *GetJobRequestType) (*GetJobResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* starts a DynectAPI session */
	SessionLogin(request *SessionLoginRequestType) (*SessionLoginResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* ends a DynectAPI session and invalidates the token */
	SessionLogout(request *SessionLogoutRequestType) (*SessionLogoutResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* checks where session and token are still valid */
	SessionIsAlive(request *SessionIsAliveRequestType) (*SessionIsAliveResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* No operation, prevents sessions from timing out */
	SessionKeepAlive(request *SessionKeepAliveRequestType) (*SessionKeepAliveResponseType, error)
}
