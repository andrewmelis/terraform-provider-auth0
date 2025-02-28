package tenant

import (
	"github.com/auth0/go-auth0/management"
)

func flattenTenantChangePassword(changePassword *management.TenantChangePassword) []interface{} {
	if changePassword == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = changePassword.Enabled
	m["html"] = changePassword.HTML

	return []interface{}{m}
}

func flattenTenantGuardianMFAPage(mfa *management.TenantGuardianMFAPage) []interface{} {
	if mfa == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = mfa.Enabled
	m["html"] = mfa.HTML

	return []interface{}{m}
}

func flattenTenantErrorPage(errorPage *management.TenantErrorPage) []interface{} {
	if errorPage == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["html"] = errorPage.HTML
	m["show_log_link"] = errorPage.ShowLogLink
	m["url"] = errorPage.URL

	return []interface{}{m}
}

func flattenTenantFlags(flags *management.TenantFlags) []interface{} {
	if flags == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enable_client_connections"] = flags.EnableClientConnections
	m["enable_apis_section"] = flags.EnableAPIsSection
	m["enable_pipeline2"] = flags.EnablePipeline2
	m["enable_dynamic_client_registration"] = flags.EnableDynamicClientRegistration
	m["enable_custom_domain_in_emails"] = flags.EnableCustomDomainInEmails
	m["universal_login"] = flags.UniversalLogin
	m["enable_legacy_logs_search_v2"] = flags.EnableLegacyLogsSearchV2
	m["disable_clickjack_protection_headers"] = flags.DisableClickjackProtectionHeaders
	m["enable_public_signup_user_exists_error"] = flags.EnablePublicSignupUserExistsError
	m["use_scope_descriptions_for_consent"] = flags.UseScopeDescriptionsForConsent
	m["allow_legacy_delegation_grant_types"] = flags.AllowLegacyDelegationGrantTypes
	m["allow_legacy_ro_grant_types"] = flags.AllowLegacyROGrantTypes
	m["allow_legacy_tokeninfo_endpoint"] = flags.AllowLegacyTokenInfoEndpoint
	m["enable_legacy_profile"] = flags.EnableLegacyProfile
	m["enable_idtoken_api2"] = flags.EnableIDTokenAPI2
	m["no_disclose_enterprise_connections"] = flags.NoDisclosureEnterpriseConnections
	m["disable_management_api_sms_obfuscation"] = flags.DisableManagementAPISMSObfuscation
	m["enable_adfs_waad_email_verification"] = flags.EnableADFSWAADEmailVerification
	m["revoke_refresh_token_grant"] = flags.RevokeRefreshTokenGrant
	m["dashboard_log_streams_next"] = flags.DashboardLogStreams
	m["dashboard_insights_view"] = flags.DashboardInsightsView
	m["disable_fields_map_fix"] = flags.DisableFieldsMapFix
	m["mfa_show_factor_list_on_enrollment"] = flags.MFAShowFactorListOnEnrollment

	return []interface{}{m}
}

func flattenTenantUniversalLogin(universalLogin *management.TenantUniversalLogin) []interface{} {
	if universalLogin == nil {
		return nil
	}
	if universalLogin.Colors == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["colors"] = []interface{}{
		map[string]interface{}{
			"primary":         universalLogin.Colors.Primary,
			"page_background": universalLogin.Colors.PageBackground,
		},
	}

	return []interface{}{m}
}

func flattenTenantSessionCookie(sessionCookie *management.TenantSessionCookie) []interface{} {
	m := make(map[string]interface{})
	m["mode"] = sessionCookie.GetMode()

	return []interface{}{m}
}
