package standardpaths

// StandardLocation This enum describes the different locations that can be queried using methods such as writableLocation, standardLocations, and displayName.
type StandardLocation int

const (
	DesktopLocation       = iota // DesktopLocation Returns the user's desktop directory. This is a generic value. On systems with no concept of a desktop.
	DocumentsLocation            //Returns the directory containing user document files. This is a generic value. The returned path is never empty.
	FontsLocation                // Returns the directory containing user's fonts. This is a generic value. Note that installing fonts may require additional, platform-specific operations.
	ApplicationsLocation         // Returns the directory containing the user applications (either executables, application bundles, or shortcuts to them). This is a generic value. Note that installing applications may require additional, platform-specific operations. Files, folders or shortcuts in this directory are platform-specific.
	MusicLocation                // Returns the directory containing the user's music or other audio files. This is a generic value. If no directory specific for music files exists, a sensible fallback for storing user documents is returned.
	MoviesLocation               // Returns the directory containing the user's movies and videos. This is a generic value. If no directory specific for movie files exists, a sensible fallback for storing user documents is returned.
	PicturesLocation             // Returns the directory containing the user's pictures or photos. This is a generic value. If no directory specific for picture files exists, a sensible fallback for storing user documents is returned.
	TempLocation                 // Returns a directory where temporary files can be stored. The returned value might be application-specific, shared among other applications for this user, or even system-wide. The returned path is never empty.
	HomeLocation                 // Returns the user's home directory (the same as QDir::homePath()). On Unix systems, this is equal to the HOME environment variable. This value might be generic or application-specific, but the returned path is never empty.
	DataLocation                 // Returns the same value as AppLocalDataLocation. This enumeration value is deprecated. Using AppDataLocation is preferable since on Windows, the roaming path is recommended.
	CacheLocation                // Returns a directory location where user-specific non-essential (cached) data should be written. This is an application-specific directory. The returned path is never empty.
	GenericCacheLocation         // Returns a directory location where user-specific non-essential (cached) data, shared across applications, should be written. This is a generic value. Note that the returned path may be empty if the system has no concept of shared cache.
	GenericDataLocation          // Returns a directory location where persistent data shared across applications can be stored. This is a generic value. The returned path is never empty.
	RuntimeLocation              // Returns a directory location where runtime communication files should be written, like Unix local sockets. This is a generic value. The returned path may be empty on some systems.
	ConfigLocation               // Returns a directory location where user-specific configuration files should be written. This may be either a generic value or application-specific, and the returned path is never empty.
	DownloadLocation             // Returns a directory for user's downloaded files. This is a generic value. If no directory specific for downloads exists, a sensible fallback for storing user documents is returned.
	GenericConfigLocation        // Returns a directory location where user-specific configuration files shared between multiple applications should be written. This is a generic value and the returned path is never empty.
	AppDataLocation              // Returns a directory location where persistent application data can be stored. This is an application-specific directory. To obtain a path to store data to be shared with other applications, use GenericDataLocation. The returned path is never empty. On the Windows operating system, this returns the roaming path. This enum value was added in Qt 5.4.
	AppLocalDataLocation         // Returns the local settings path on the Windows operating system. On all other platforms, it returns the same value as AppDataLocation. This enum value was added in Qt 5.4.
	AppConfigLocation            // Returns a directory location where user-specific configuration files should be written. This is an application-specific directory, and the returned path is never empty. This enum value was added in Qt 5.5.
)
