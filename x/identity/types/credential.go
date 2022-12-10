package types

// p "github.com/duo-labs/webauthn/protocol"
// "github.com/duo-labs/webauthn/webauthn"


// VerifyCounter
// Step 17 of §7.2. about verifying attestation. If the signature counter value authData.signCount
// is nonzero or the value stored in conjunction with credential’s id attribute is nonzero, then
// run the following sub-step:
//
//  If the signature counter value authData.signCount is
//
//  → Greater than the signature counter value stored in conjunction with credential’s id attribute.
//  Update the stored signature counter value, associated with credential’s id attribute, to be the value of
//  authData.signCount.
//
//  → Less than or equal to the signature counter value stored in conjunction with credential’s id attribute.
//  This is a signal that the authenticator may be cloned, see CloneWarning above for more information.
func (a *WebauthnAuthenticator) UpdateCounter(authDataCount uint32) {
	if authDataCount <= a.SignCount && (authDataCount != 0 || a.SignCount != 0) {
		a.CloneWarning = true
		return
	}
	a.SignCount = authDataCount
}
