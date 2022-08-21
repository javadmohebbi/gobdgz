package gobdgz

type PolicyItems struct {
	Name       string `json:"name"`
	UISettings struct {
		General struct {
			Display struct {
				SilentMode struct {
					Enable               bool `json:"enable"`
					IconNotificationArea bool `json:"iconNotificationArea"`
					ShowPopups           struct {
						Enable bool `json:"enable"`
					} `json:"showPopups"`
					DisplayAlertsPopups bool `json:"displayAlertsPopups"`
					ShowBrowserToolbar  bool `json:"showBrowserToolbar"`
					IssuesVisibility    struct {
						Enable                           bool `json:"enable"`
						Profile                          int  `json:"profile"`
						General                          bool `json:"general"`
						Antimalware                      bool `json:"antimalware"`
						Firewall                         bool `json:"firewall"`
						ContentControl                   bool `json:"contentControl"`
						Update                           bool `json:"update"`
						InstallationRestartNotifications int  `json:"installationRestartNotifications"`
						OnAccessNotifications            int  `json:"onAccessNotifications"`
						OnDemandNotifications            int  `json:"onDemandNotifications"`
						DissinfectionRestartNotification int  `json:"dissinfectionRestartNotification"`
						UpdateRestartNotifications       int  `json:"updateRestartNotifications"`
						CloudConnectionNotifications     int  `json:"cloudConnectionNotifications"`
						ShowAfter                        int  `json:"showAfter"`
						Introspection                    bool `json:"introspection"`
					} `json:"issuesVisibility"`
					EndpointRestartPopup struct {
						Enable          bool `json:"enable"`
						Update          bool `json:"update"`
						PatchManagement bool `json:"patchManagement"`
					} `json:"endpointRestartPopup"`
				} `json:"silentMode"`
				SupportInformation struct {
					Website string `json:"website"`
					Email   string `json:"email"`
					Phone   string `json:"phone"`
				} `json:"supportInformation"`
			} `json:"display"`
			Advanced struct {
				Settings struct {
					ScanSsl struct {
						Enabled   bool `json:"enabled"`
						Protocols struct {
							Incoming struct {
								Rdp bool `json:"rdp"`
							} `json:"incoming"`
							Outgoing struct {
								HTTP bool `json:"http"`
							} `json:"outgoing"`
						} `json:"protocols"`
					} `json:"scanSsl"`
					BrowserSearchAdvisor bool `json:"browserSearchAdvisor"`
					RemoveEvents         struct {
						Settings struct {
							Days int `json:"days"`
						} `json:"settings"`
					} `json:"removeEvents"`
					SubmitReports              bool `json:"submitReports"`
					SubmitSuspicious           bool `json:"submitSuspicious"`
					SendTelemetry              bool `json:"sendTelemetry"`
					UseGlobalProtectiveNetwork bool `json:"useGlobalProtectiveNetwork"`
					ReportActiveSessions       bool `json:"reportActiveSessions"`
				} `json:"settings"`
				PasswordConfig struct {
					Profile int    `json:"profile"`
					Value   string `json:"value"`
				} `json:"passwordConfig"`
				PowerUser struct {
					Enabled  bool   `json:"enabled"`
					Password string `json:"password"`
				} `json:"powerUser"`
			} `json:"advanced"`
			Communication struct {
				EcsAssignments []interface{} `json:"ecsAssignments"`
				EcsProxy       struct {
					Profile int `json:"profile"`
				} `json:"ecsProxy"`
				CloudServicesProxy struct {
					Profile int `json:"profile"`
				} `json:"cloudServicesProxy"`
			} `json:"communication"`
			Update struct {
				Settings struct {
					ProductUpdateScheduler struct {
						Enabled        bool `json:"enabled"`
						Occurrence     int  `json:"occurrence"`
						UpdateInterval int  `json:"updateInterval"`
						WeekDays       struct {
							Days        []int `json:"days"`
							StartHour   int   `json:"startHour"`
							StartMinute int   `json:"startMinute"`
							EndHour     int   `json:"endHour"`
							EndMinute   int   `json:"endMinute"`
						} `json:"weekDays"`
						PostponeReboot        bool `json:"postponeReboot"`
						RebootAfterInstalling struct {
							Enable   bool `json:"enable"`
							Settings struct {
								Day     int `json:"day"`
								Hour    int `json:"hour"`
								Minutes int `json:"minutes"`
							} `json:"settings"`
						} `json:"rebootAfterInstalling"`
						UpdateLinuxEdrUsingProductUpdate bool `json:"updateLinuxEdrUsingProductUpdate"`
					} `json:"productUpdateScheduler"`
					SignatureUpdateScheduler struct {
						Enabled        bool `json:"enabled"`
						Occurrence     int  `json:"occurrence"`
						UpdateInterval int  `json:"updateInterval"`
						WeekDays       struct {
							Days        []int `json:"days"`
							StartHour   int   `json:"startHour"`
							StartMinute int   `json:"startMinute"`
							EndHour     int   `json:"endHour"`
							EndMinute   int   `json:"endMinute"`
						} `json:"weekDays"`
					} `json:"signatureUpdateScheduler"`
					UpdateRing int `json:"updateRing"`
					Proxy      struct {
						Enable   bool `json:"enable"`
						Profile  int  `json:"profile"`
						Settings struct {
							Server   string `json:"server"`
							Port     int    `json:"port"`
							Username string `json:"username"`
							Password string `json:"password"`
						} `json:"settings"`
					} `json:"proxy"`
					DefaultLocation bool `json:"defaultLocation"`
				} `json:"settings"`
				UpdateLocations []struct {
					Server   string `json:"server"`
					UseProxy bool   `json:"useProxy"`
				} `json:"updateLocations"`
			} `json:"update"`
			SecurityTelemetry struct {
				Enabled                    bool   `json:"enabled"`
				Siem                       int    `json:"siem"`
				Protocol                   int    `json:"protocol"`
				Format                     int    `json:"format"`
				URL                        string `json:"url"`
				AllowSelfSignedCertificate bool   `json:"allowSelfSignedCertificate"`
				Key                        string `json:"key"`
				Proxy                      struct {
					Profile int `json:"profile"`
				} `json:"proxy"`
				EventTypes struct {
					CreateFile        bool `json:"createFile"`
					CreateProcess     bool `json:"createProcess"`
					DeleteFile        bool `json:"deleteFile"`
					Logon             bool `json:"logon"`
					Logout            bool `json:"logout"`
					ModifyFile        bool `json:"modifyFile"`
					MoveFile          bool `json:"moveFile"`
					NetworkConnection bool `json:"networkConnection"`
					ReadFromFile      bool `json:"readFromFile"`
					RegCreateKey      bool `json:"regCreateKey"`
					RegDeleteKey      bool `json:"regDeleteKey"`
					RegDeleteValue    bool `json:"regDeleteValue"`
					RegModifyValue    bool `json:"regModifyValue"`
					TerminateProcess  bool `json:"terminateProcess"`
				} `json:"eventTypes"`
			} `json:"securityTelemetry"`
			AllowChangeByOtherUsers bool `json:"allowChangeByOtherUsers"`
		} `json:"general"`
		Antimalware struct {
			OnAccess struct {
				OnAccessScanning struct {
					Enable   bool `json:"enable"`
					Profile  int  `json:"profile"`
					Settings struct {
						General struct {
							FileTypes struct {
								ScanLocalFile struct {
									Enable   bool `json:"enable"`
									Settings struct {
										FileType   int    `json:"fileType"`
										Extensions string `json:"extensions"`
									} `json:"settings"`
								} `json:"scanLocalFile"`
								ScanNetworkFile struct {
									Enable   bool `json:"enable"`
									Settings struct {
										FileType   int    `json:"fileType"`
										Extensions string `json:"extensions"`
									} `json:"settings"`
								} `json:"scanNetworkFile"`
								LimitFile struct {
									Enable bool `json:"enable"`
									Size   int  `json:"size"`
								} `json:"limitFile"`
							} `json:"fileTypes"`
							Archives struct {
								Enable              bool `json:"enable"`
								LimitArchiveSize    int  `json:"limitArchiveSize"`
								MaximumArchiveDepth int  `json:"maximumArchiveDepth"`
							} `json:"archives"`
							Miscellaneous struct {
								ScanBootSectors        bool `json:"scanBootSectors"`
								ScanProcessMemory      bool `json:"scanProcessMemory"`
								ScanOnlyNewChangeFiles bool `json:"scanOnlyNewChangeFiles"`
								ScanForKeyloggers      bool `json:"scanForKeyloggers"`
								ScanPUA                bool `json:"scanPUA"`
								DeferredScanning       bool `json:"deferredScanning"`
							} `json:"miscellaneous"`
							ScanAction struct {
								InfectedFiles struct {
									Action int `json:"action"`
									Then   int `json:"then"`
								} `json:"infectedFiles"`
								SuspectFiles struct {
									Action int `json:"action"`
									Then   int `json:"then"`
								} `json:"suspectFiles"`
							} `json:"scanAction"`
						} `json:"general"`
						Advanced struct {
							OnAccessUnix      bool     `json:"onAccessUnix"`
							OnAccessUnixPaths []string `json:"onAccessUnixPaths"`
						} `json:"advanced"`
					} `json:"settings"`
				} `json:"onAccessScanning"`
				RansomwareProtection struct {
					Enable bool `json:"enable"`
				} `json:"ransomwareProtection"`
			} `json:"onAccess"`
			OnExecute struct {
				VirusControl struct {
					Enable        bool `json:"enable"`
					Profile       int  `json:"profile"`
					DefaultAction int  `json:"defaultAction"`
				} `json:"virusControl"`
				CommandLineScanning struct {
					FilelessAttackProtection bool `json:"filelessAttackProtection"`
					Enable                   bool `json:"enable"`
					Amsi                     struct {
						Enable                   bool `json:"enable"`
						ReportDetectionsToCaller bool `json:"reportDetectionsToCaller"`
					} `json:"amsi"`
				} `json:"commandLineScanning"`
				Theta struct {
					Enabled bool `json:"enabled"`
				} `json:"theta"`
				RansomwareRemediation struct {
					Enable      bool `json:"enable"`
					LocalAttack struct {
						Enable bool `json:"enable"`
					} `json:"localAttack"`
					RemoteAttack struct {
						Enable bool `json:"enable"`
					} `json:"remoteAttack"`
					Restore struct {
						Mode int `json:"mode"`
					} `json:"restore"`
				} `json:"ransomwareRemediation"`
			} `json:"onExecute"`
			DynamicThreatDefense struct {
				Enable                          bool `json:"enable"`
				EnableGlobalReportingForFiles   bool `json:"enableGlobalReportingForFiles"`
				EnableGlobalReportingForNetwork bool `json:"enableGlobalReportingForNetwork"`
				ActionforLocal                  int  `json:"actionforLocal"`
				ActionforNetwork                int  `json:"actionforNetwork"`
				TargetedAttack                  struct {
					Enable         bool `json:"enable"`
					DetectionLevel int  `json:"detectionLevel"`
				} `json:"targetedAttack"`
				SuspiciousFilesAndNetworkTraffic struct {
					Enable         bool `json:"enable"`
					DetectionLevel int  `json:"detectionLevel"`
				} `json:"suspiciousFilesAndNetworkTraffic"`
				Exploits struct {
					Enable         bool `json:"enable"`
					DetectionLevel int  `json:"detectionLevel"`
				} `json:"exploits"`
				Ransomware struct {
					Enable         bool `json:"enable"`
					DetectionLevel int  `json:"detectionLevel"`
				} `json:"ransomware"`
				Grayware struct {
					Enable         bool `json:"enable"`
					DetectionLevel int  `json:"detectionLevel"`
				} `json:"grayware"`
			} `json:"dynamicThreatDefense"`
			OnDemand struct {
				ScanTask       []interface{} `json:"scanTask"`
				DeviceScanning struct {
					Enable   bool `json:"enable"`
					Settings struct {
						AutomaticallyScanCdDvd bool `json:"automaticallyScanCdDvd"`
						AutomaticallyScanUsb   bool `json:"automaticallyScanUsb"`
						ScanDevices            struct {
							Enable bool `json:"enable"`
							Maxim  int  `json:"maxim"`
						} `json:"scanDevices"`
					} `json:"settings"`
					ScanProfile  int `json:"scanProfile"`
					ScanSettings struct {
						FileTypes  int           `json:"fileTypes"`
						Extensions []interface{} `json:"extensions"`
						Archives   struct {
							Enable   bool `json:"enable"`
							Settings struct {
								LimitArchiveSize    int `json:"limitArchiveSize"`
								MaximumArchiveDepth int `json:"maximumArchiveDepth"`
							} `json:"settings"`
							ScanEmail bool `json:"scanEmail"`
						} `json:"archives"`
						Miscellaneous struct {
							ScanBoot         bool `json:"scanBoot"`
							ScanRegistry     bool `json:"scanRegistry"`
							ScanRootkits     bool `json:"scanRootkits"`
							IgnoreKeyloggers bool `json:"ignoreKeyloggers"`
							ScanMemory       bool `json:"scanMemory"`
							ScanCookie       bool `json:"scanCookie"`
							ScanNewChanged   bool `json:"scanNewChanged"`
							ScanPUA          bool `json:"scanPUA"`
							ScanNetworkFiles bool `json:"scanNetworkFiles"`
						} `json:"miscellaneous"`
						Action struct {
							WhenInfected struct {
								Action int `json:"action"`
								Then   int `json:"then"`
							} `json:"whenInfected"`
							WhenSuspect struct {
								Action int `json:"action"`
								Then   int `json:"then"`
							} `json:"whenSuspect"`
							WhenRootKitAction int `json:"whenRootKitAction"`
						} `json:"action"`
					} `json:"scanSettings"`
				} `json:"deviceScanning"`
				ContextualScan struct {
					ScanProfile  int `json:"scanProfile"`
					ScanSettings struct {
						Miscellaneous struct {
							ScanBoot         bool `json:"scanBoot"`
							ScanRegistry     bool `json:"scanRegistry"`
							ScanRootkits     bool `json:"scanRootkits"`
							IgnoreKeyloggers bool `json:"ignoreKeyloggers"`
							ScanMemory       bool `json:"scanMemory"`
							ScanCookie       bool `json:"scanCookie"`
							ScanNewChanged   bool `json:"scanNewChanged"`
							ScanPUA          bool `json:"scanPUA"`
							ScanNetworkFiles bool `json:"scanNetworkFiles"`
						} `json:"miscellaneous"`
						Action struct {
							WhenInfected struct {
								Action int `json:"action"`
								Then   int `json:"then"`
							} `json:"whenInfected"`
							WhenSuspect struct {
								Action int `json:"action"`
								Then   int `json:"then"`
							} `json:"whenSuspect"`
							WhenRootKitAction int `json:"whenRootKitAction"`
						} `json:"action"`
						Archives struct {
							Enable   bool `json:"enable"`
							Settings struct {
								LimitArchiveSize    int `json:"limitArchiveSize"`
								MaximumArchiveDepth int `json:"maximumArchiveDepth"`
							} `json:"settings"`
							ScanEmail bool `json:"scanEmail"`
						} `json:"archives"`
					} `json:"scanSettings"`
				} `json:"contextualScan"`
			} `json:"onDemand"`
			AntiExploit struct {
				Enable                 bool `json:"enable"`
				PredefinedApplications []struct {
					Details struct {
						ApplicationName string   `json:"applicationName"`
						Status          int      `json:"status"`
						ProcessName     []string `json:"processName"`
					} `json:"details"`
					ExploitDetectionTechniques []struct {
						Enable bool `json:"enable"`
						Name   int  `json:"name"`
						Action int  `json:"action"`
					} `json:"exploitDetectionTechniques"`
				} `json:"predefinedApplications"`
				CustomApplications   []interface{} `json:"customApplications"`
				SystemWideDetections struct {
					PrivilegeEscalationStatus                 bool `json:"privilegeEscalationStatus"`
					PrivilegeEscalation                       int  `json:"privilegeEscalation"`
					LsassMemoryAccessFromUnknownProcess       int  `json:"lsassMemoryAccessFromUnknownProcess"`
					LsassMemoryAccessFromUnknownProcessStatus bool `json:"lsassMemoryAccessFromUnknownProcessStatus"`
				} `json:"systemWideDetections"`
				Linux struct {
					SystemWideDetections struct {
						CredentialsStatus bool `json:"credentialsStatus"`
						Credentials       int  `json:"credentials"`
						PtraceStatus      bool `json:"ptraceStatus"`
						Ptrace            int  `json:"ptrace"`
						NamespaceStatus   bool `json:"namespaceStatus"`
						Namespace         int  `json:"namespace"`
						CorruptionStatus  bool `json:"corruptionStatus"`
						Corruption        int  `json:"corruption"`
						PermissionsStatus bool `json:"permissionsStatus"`
						Permissions       int  `json:"permissions"`
					} `json:"systemWideDetections"`
				} `json:"linux"`
			} `json:"antiExploit"`
			Settings struct {
				ActivateExclusions struct {
					Enable                       bool          `json:"enable"`
					ExclusionsItems              []interface{} `json:"exclusionsItems"`
					UseVendorExclusionLists      bool          `json:"useVendorExclusionLists"`
					SelectedVendorExclusionLists []interface{} `json:"selectedVendorExclusionLists"`
					ExclusionsLists              []interface{} `json:"exclusionsLists"`
					UseExclusionLists            bool          `json:"useExclusionLists"`
				} `json:"activateExclusions"`
				UseBuiltInExclusions bool `json:"useBuiltInExclusions"`
				DeleteFilesOlderThan int  `json:"deleteFilesOlderThan"`
				SubmitQuarantined    struct {
					Enable bool `json:"enable"`
					Hours  int  `json:"hours"`
				} `json:"submitQuarantined"`
				RescanQuarantine                     bool `json:"rescanQuarantine"`
				CopyFilesToQuarantineBeforeDisinfect bool `json:"copyFilesToQuarantineBeforeDisinfect"`
				AllowUserQuarantineActions           bool `json:"allowUserQuarantineActions"`
				CentralizedQuarantine                struct {
					AutoSubmitToSandboxAnalyzer bool   `json:"autoSubmitToSandboxAnalyzer"`
					Enable                      bool   `json:"enable"`
					ArchivePassword             string `json:"archivePassword"`
					Location                    struct {
						Path     string `json:"path"`
						Username string `json:"username"`
						Password string `json:"password"`
					} `json:"location"`
				} `json:"centralizedQuarantine"`
			} `json:"settings"`
			SecurityServers struct {
				AffinityRule        bool          `json:"affinityRule"`
				SvaMpAffinityRules  bool          `json:"svaMpAffinityRules"`
				SvaIps              []interface{} `json:"svaIps"`
				UseSSL              bool          `json:"useSSL"`
				EnableOnDemandSlots bool          `json:"enableOnDemandSlots"`
				SvaProxy            struct {
					Profile int `json:"profile"`
				} `json:"svaProxy"`
				ServerOrder string `json:"serverOrder"`
			} `json:"securityServers"`
		} `json:"antimalware"`
		NetworkSandboxing struct {
			Enabled            bool `json:"enabled"`
			AnalysisMode       int  `json:"analysisMode"`
			RemediationActions struct {
				DefaultAction  int `json:"defaultAction"`
				FallbackAction int `json:"fallbackAction"`
			} `json:"remediationActions"`
			Settings struct {
				IsCloudSandbox    bool        `json:"isCloudSandbox"`
				IP                string      `json:"ip"`
				Hostname          string      `json:"hostname"`
				EndpointID        interface{} `json:"endpointId"`
				DetonationProfile int         `json:"detonationProfile"`
			} `json:"settings"`
			Proxy struct {
				Enabled  bool `json:"enabled"`
				Settings struct {
					Server   string `json:"server"`
					Port     int    `json:"port"`
					Username string `json:"username"`
					Password string `json:"password"`
				} `json:"settings"`
			} `json:"proxy"`
			ContentPrefiltering struct {
				Categories struct {
					Applications struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"applications"`
					Documents struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"documents"`
					Scripts struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"scripts"`
					Archives struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"archives"`
					Emails struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"emails"`
				} `json:"categories"`
				ExcludeExtensions []interface{} `json:"excludeExtensions"`
				SizeFilter        struct {
					Enabled       bool `json:"enabled"`
					FileMinSizeKB int  `json:"fileMinSizeKB"`
					FileMaxSizeMB int  `json:"fileMaxSizeMB"`
				} `json:"sizeFilter"`
			} `json:"contentPrefiltering"`
		} `json:"networkSandboxing"`
		SandboxManager struct {
			DataRetention  int  `json:"dataRetention"`
			PersistSamples bool `json:"persistSamples"`
		} `json:"sandboxManager"`
		Ghostr struct {
			Enabled                 bool `json:"enabled"`
			SandboxAnalyzerSettings struct {
				IP                string      `json:"ip"`
				Hostname          string      `json:"hostname"`
				Fqdn              string      `json:"fqdn"`
				EndpointID        interface{} `json:"endpointId"`
				DetonationProfile int         `json:"detonationProfile"`
			} `json:"sandboxAnalyzerSettings"`
			Proxy struct {
				Enabled  bool `json:"enabled"`
				Settings struct {
					Server   string `json:"server"`
					Port     int    `json:"port"`
					Username string `json:"username"`
					Password string `json:"password"`
				} `json:"settings"`
			} `json:"proxy"`
			ContentPrefiltering struct {
				Categories struct {
					Applications struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"applications"`
					Documents struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"documents"`
					Scripts struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"scripts"`
					Archives struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"archives"`
					Emails struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"emails"`
				} `json:"categories"`
				ExcludeExtensions []interface{} `json:"excludeExtensions"`
				SizeFilter        struct {
					Enabled       bool `json:"enabled"`
					FileMinSizeKB int  `json:"fileMinSizeKB"`
					FileMaxSizeMB int  `json:"fileMaxSizeMB"`
				} `json:"sizeFilter"`
			} `json:"contentPrefiltering"`
		} `json:"ghostr"`
		IcapSensor struct {
			Enabled                 bool `json:"enabled"`
			SandboxAnalyzerSettings struct {
				IP                string      `json:"ip"`
				Hostname          string      `json:"hostname"`
				Fqdn              string      `json:"fqdn"`
				EndpointID        interface{} `json:"endpointId"`
				DetonationProfile int         `json:"detonationProfile"`
			} `json:"sandboxAnalyzerSettings"`
			Proxy struct {
				Enabled  bool `json:"enabled"`
				Settings struct {
					Server   string `json:"server"`
					Port     int    `json:"port"`
					Username string `json:"username"`
					Password string `json:"password"`
				} `json:"settings"`
			} `json:"proxy"`
			ContentPrefiltering struct {
				Categories struct {
					Applications struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"applications"`
					Documents struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"documents"`
					Scripts struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"scripts"`
					Archives struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"archives"`
					Emails struct {
						Enabled bool `json:"enabled"`
						Level   int  `json:"level"`
					} `json:"emails"`
				} `json:"categories"`
				ExcludeExtensions []interface{} `json:"excludeExtensions"`
				SizeFilter        struct {
					Enabled       bool `json:"enabled"`
					FileMinSizeKB int  `json:"fileMinSizeKB"`
					FileMaxSizeMB int  `json:"fileMaxSizeMB"`
				} `json:"sizeFilter"`
				SubmitFilesMarkedAsMalware bool `json:"submitFilesMarkedAsMalware"`
			} `json:"contentPrefiltering"`
		} `json:"icapSensor"`
		Firewall struct {
			Settings struct {
				Enable                         bool `json:"enable"`
				BlockPortScan                  bool `json:"blockPortScan"`
				AllowInternetConnectionSharing bool `json:"allowInternetConnectionSharing"`
				MonitorWiFiConnections         bool `json:"monitorWiFiConnections"`
				LogVerbosity                   struct {
					Enable bool `json:"enable"`
					Level  int  `json:"level"`
				} `json:"logVerbosity"`
				InstructionsDetectionSystem struct {
					Enable  bool `json:"enable"`
					Profile int  `json:"profile"`
				} `json:"instructionsDetectionSystem"`
			} `json:"settings"`
			Network struct {
				Adapters []struct {
					Type        string `json:"type"`
					NetworkType int    `json:"networkType"`
					StealthMode int    `json:"stealthMode"`
				} `json:"adapters"`
			} `json:"network"`
			Advanced struct {
				ProtectionLevel                int  `json:"protectionLevel"`
				CreateAgresiveRules            bool `json:"createAgresiveRules"`
				CreateRulesForAppsBlockedByIDS bool `json:"createRulesForAppsBlockedByIDS"`
				MonitorProcessChanges          bool `json:"monitorProcessChanges"`
				IgnoreSignedProcesses          bool `json:"ignoreSignedProcesses"`
				Rule                           []struct {
					RuleID      string `json:"ruleId"`
					DefaultRule int    `json:"defaultRule"`
					Type        int    `json:"type"`
					RuleType    int    `json:"ruleType"`
					Details     struct {
						Name           string `json:"name"`
						ApplictionPath string `json:"applictionPath"`
						CommandLine    string `json:"commandLine"`
						ApplicationMd5 string `json:"applicationMd5"`
					} `json:"details"`
					Settings struct {
						LocalAddress struct {
							Any       bool   `json:"any"`
							PortRange string `json:"portRange"`
						} `json:"localAddress"`
						RemoteAddress struct {
							Any       bool   `json:"any"`
							PortRange string `json:"portRange"`
						} `json:"remoteAddress"`
						DirectlyConnected struct {
							Enable bool `json:"enable"`
						} `json:"directlyConnected"`
						Protocol       int    `json:"protocol"`
						CustomProtocol string `json:"customProtocol"`
						Direction      int    `json:"direction"`
						IPVersion      int    `json:"ipVersion"`
					} `json:"settings"`
					Permission struct {
						Home          bool `json:"home"`
						Public        bool `json:"public"`
						SetPermission int  `json:"setPermission"`
					} `json:"permission"`
				} `json:"rule"`
			} `json:"advanced"`
		} `json:"firewall"`
		ContentControl struct {
			WebAccess struct {
				Enable        bool          `json:"enable"`
				ProfileType   int           `json:"profileType"`
				UseExceptions bool          `json:"useExceptions"`
				WebRules      []interface{} `json:"webRules"`
				Scheduler     []interface{} `json:"scheduler"`
			} `json:"webAccess"`
			WebCategoriesFilter struct {
				Enable               bool `json:"enable"`
				ProfileType          int  `json:"profileType"`
				TreatAsExceptions    bool `json:"treatAsExceptions"`
				EnableDetailedAlerts bool `json:"enableDetailedAlerts"`
			} `json:"webCategoriesFilter"`
			Antiphishing struct {
				Enable        bool `json:"enable"`
				DefaultAction int  `json:"defaultAction"`
				Settings      struct {
					ProtectionAgainstFraud    bool `json:"protectionAgainstFraud"`
					ProtectionAgainstPhishing bool `json:"protectionAgainstPhishing"`
				} `json:"settings"`
			} `json:"antiphishing"`
			Application struct {
				Enable bool          `json:"enable"`
				Rules  []interface{} `json:"rules"`
			} `json:"application"`
			DataProtection struct {
				Enable     bool          `json:"enable"`
				Rules      []interface{} `json:"rules"`
				Exclusions []interface{} `json:"exclusions"`
			} `json:"dataProtection"`
			Traffic struct {
				Enable           bool `json:"enable"`
				EnableExclusions bool `json:"enableExclusions"`
				TrafficScan      struct {
					IncomingEmails bool `json:"incomingEmails"`
					OutgoingEmails bool `json:"outgoingEmails"`
					WebTraffic     bool `json:"webTraffic"`
					EmailTraffic   bool `json:"emailTraffic"`
				} `json:"trafficScan"`
				Exclusions []interface{} `json:"exclusions"`
			} `json:"traffic"`
			NetworkMonitor struct {
				Enable           bool `json:"enable"`
				AttackTechniques []struct {
					Enable bool `json:"enable"`
					Name   int  `json:"name"`
					Action int  `json:"action"`
				} `json:"attackTechniques"`
			} `json:"networkMonitor"`
		} `json:"contentControl"`
		ApplicationControl struct {
			General struct {
				Enable     bool `json:"enable"`
				ReportOnly bool `json:"reportOnly"`
			} `json:"general"`
			StartRules []interface{} `json:"startRules"`
		} `json:"applicationControl"`
		Exchange struct {
			UserGroups []interface{} `json:"userGroups"`
			BlackList  struct {
				Enable bool          `json:"enable"`
				List   []interface{} `json:"list"`
			} `json:"blackList"`
			Quarantine struct {
				DeleteFilesOlderThan int `json:"deleteFilesOlderThan"`
			} `json:"quarantine"`
			Antimalware struct {
				Enable   bool `json:"enable"`
				Settings struct {
					AppendFooter               bool   `json:"appendFooter"`
					FooterText                 string `json:"footerText"`
					InfectedText               string `json:"infectedText"`
					QuarantinedText            string `json:"quarantinedText"`
					UnscannableInfectedText    string `json:"unscannableInfectedText"`
					UnscannableQuarantinedText string `json:"unscannableQuarantinedText"`
				} `json:"settings"`
				Rules []struct {
					Name    string `json:"name"`
					Active  bool   `json:"active"`
					Default bool   `json:"default"`
					Scope   struct {
						ApplyTo int `json:"applyTo"`
						From    int `json:"from"`
						To      int `json:"to"`
					} `json:"scope"`
					Settings struct {
						ScanMode            int           `json:"scanMode"`
						Extensions          []interface{} `json:"extensions"`
						MaximumSizeEnabled  bool          `json:"maximumSizeEnabled"`
						MaximumSize         int           `json:"maximumSize"`
						MaximumDepthEnabled bool          `json:"maximumDepthEnabled"`
						MaximumDepth        int           `json:"maximumDepth"`
						ScanPUA             bool          `json:"scanPUA"`
					} `json:"settings"`
					Actions struct {
						Infected struct {
							Action int `json:"action"`
							Then   int `json:"then"`
						} `json:"infected"`
						Suspected struct {
							Action int `json:"action"`
							Then   int `json:"then"`
						} `json:"suspected"`
						Unscannable struct {
							Action int `json:"action"`
							Then   int `json:"then"`
						} `json:"unscannable"`
						StopIfMatched bool `json:"stopIfMatched"`
					} `json:"actions"`
				} `json:"rules"`
				Exclusions []interface{} `json:"exclusions"`
				ScanTasks  []interface{} `json:"scanTasks"`
			} `json:"antimalware"`
			Antispoofing struct {
				Enable       bool          `json:"enable"`
				DomainIPList []interface{} `json:"domainIpList"`
			} `json:"antispoofing"`
			Antispam struct {
				Enable   bool `json:"enable"`
				Settings struct {
					RblDNSAddress   string        `json:"rblDnsAddress"`
					RblDNSMsTimeout int           `json:"rblDnsMsTimeout"`
					RblServers      []interface{} `json:"rblServers"`
				} `json:"settings"`
				WhiteList struct {
					Enable bool          `json:"enable"`
					List   []interface{} `json:"list"`
				} `json:"whiteList"`
				Rules []struct {
					Name    string `json:"name"`
					Default bool   `json:"default"`
					Active  bool   `json:"active"`
					Scope   struct {
						ApplyTo int `json:"applyTo"`
						From    int `json:"from"`
						To      int `json:"to"`
					} `json:"scope"`
					Aggressivity int `json:"aggressivity"`
					Filters      struct {
						Asian       bool `json:"asian"`
						Cyrillic    bool `json:"cyrillic"`
						FapMaterial bool `json:"fapMaterial"`
						URL         bool `json:"url"`
						Rbl         bool `json:"rbl"`
						Cloud       bool `json:"cloud"`
						Heuristic   bool `json:"heuristic"`
					} `json:"filters"`
					CheckAuthConnections bool `json:"checkAuthConnections"`
					Actions              struct {
						Type              int    `json:"type"`
						ExchangeSCL       bool   `json:"exchangeSCL"`
						ModifySubject     bool   `json:"modifySubject"`
						ModifySubjectText string `json:"modifySubjectText"`
						AppendHeader      bool   `json:"appendHeader"`
						HeaderName        string `json:"headerName"`
						HeaderValue       string `json:"headerValue"`
						SaveMailToDisk    bool   `json:"saveMailToDisk"`
						ArchiveToAccount  bool   `json:"archiveToAccount"`
						StopIfMatched     bool   `json:"stopIfMatched"`
					} `json:"actions"`
				} `json:"rules"`
			} `json:"antispam"`
			ContentControl struct {
				AttachmentFiltering struct {
					Enable   bool `json:"enable"`
					Settings struct {
						ReplacementText string `json:"replacementText"`
					} `json:"settings"`
					Exclusions struct {
						TrustedSenders                       []interface{} `json:"trustedSenders"`
						TrustedRecipients                    []interface{} `json:"trustedRecipients"`
						ExcludeOnlyIfAllRecipientsAreTrusted bool          `json:"excludeOnlyIfAllRecipientsAreTrusted"`
					} `json:"exclusions"`
					Rules []interface{} `json:"rules"`
				} `json:"attachmentFiltering"`
				ContentFiltering struct {
					Enable     bool `json:"enable"`
					Exclusions struct {
						TrustedSenders                       []interface{} `json:"trustedSenders"`
						TrustedRecipients                    []interface{} `json:"trustedRecipients"`
						ExcludeOnlyIfAllRecipientsAreTrusted bool          `json:"excludeOnlyIfAllRecipientsAreTrusted"`
					} `json:"exclusions"`
					Rules []interface{} `json:"rules"`
				} `json:"contentFiltering"`
			} `json:"contentControl"`
		} `json:"exchange"`
		DeviceControl struct {
			DataLossPrevention struct {
				Enable bool `json:"enable"`
				Rules  []struct {
					DeviceClass int           `json:"deviceClass"`
					RuleName    string        `json:"ruleName"`
					ProductIds  []interface{} `json:"productIds"`
					DeviceIds   []interface{} `json:"deviceIds"`
					Permissions struct {
						ALL     int `json:"ALL"`
						PCI     int `json:"PCI"`
						PCMCIA  int `json:"PCMCIA"`
						USB     int `json:"USB"`
						UNKNOWN int `json:"UNKNOWN"`
					} `json:"permissions"`
				} `json:"rules"`
			} `json:"dataLossPrevention"`
			DataLossPreventionExceptions struct {
				Enable bool          `json:"enable"`
				Rules  []interface{} `json:"rules"`
			} `json:"dataLossPreventionExceptions"`
		} `json:"deviceControl"`
		Relay struct {
			Communication struct {
				CloudServicesProxy struct {
					Profile int `json:"profile"`
				} `json:"cloudServicesProxy"`
				ApplianceProxy struct {
					Profile int `json:"profile"`
				} `json:"applianceProxy"`
			} `json:"communication"`
			Update struct {
				UpdateInterval struct {
					Enable bool `json:"enable"`
					Hours  int  `json:"hours"`
				} `json:"updateInterval"`
				DownloadFolder  string `json:"downloadFolder"`
				UpdateLocations struct {
					Enable    bool `json:"enable"`
					Locations []struct {
						Server   string `json:"server"`
						UseProxy bool   `json:"useProxy"`
					} `json:"locations"`
				} `json:"updateLocations"`
			} `json:"update"`
		} `json:"relay"`
		Nsx struct {
			Enabled bool        `json:"enabled"`
			Name    interface{} `json:"name"`
			Default bool        `json:"default"`
		} `json:"nsx"`
		Encryption struct {
			Enabled       bool `json:"enabled"`
			Mode          int  `json:"mode"`
			EncryptPolicy struct {
				Tpm struct {
					ShouldAskForPassword bool `json:"shouldAskForPassword"`
				} `json:"tpm"`
			} `json:"encryptPolicy"`
			Exclusions struct {
				Enabled bool          `json:"enabled"`
				Items   []interface{} `json:"items"`
			} `json:"exclusions"`
		} `json:"encryption"`
		Hvi struct {
			UserSpaceMemoryIntrospection struct {
				Enabled            bool `json:"enabled"`
				ApplicationCrashes bool `json:"applicationCrashes"`
				ConnectionEvents   bool `json:"connectionEvents"`
				Rules              []struct {
					RuleName          string   `json:"ruleName"`
					Type              int      `json:"type"`
					Processes         []string `json:"processes"`
					MonitoringMode    int      `json:"monitoringMode"`
					MonitoringOptions struct {
						DllHooks              bool `json:"dllHooks"`
						ExeUnpackAttemtps     bool `json:"exeUnpackAttemtps"`
						ProcessRemoteWrites   bool `json:"processRemoteWrites"`
						Exploits              bool `json:"exploits"`
						WinSockHooking        bool `json:"winSockHooking"`
						DoubleAgentPrevention bool `json:"doubleAgentPrevention"`
					} `json:"monitoringOptions"`
					Actions struct {
						PrimaryAction           int `json:"primaryAction"`
						RemediationAction       int `json:"remediationAction"`
						BackupRemediationAction int `json:"backupRemediationAction"`
					} `json:"actions"`
				} `json:"rules"`
			} `json:"userSpaceMemoryIntrospection"`
			KernelSpaceMemoryIntrospection struct {
				Enabled           bool `json:"enabled"`
				MonitoringMode    int  `json:"monitoringMode"`
				MonitoringOptions struct {
					ControlRegisters       bool `json:"controlRegisters"`
					ModelSpecificRegisters bool `json:"modelSpecificRegisters"`
					IdtOrGdtIntegrity      bool `json:"idtOrGdtIntegrity"`
					AntimalwareDrivers     bool `json:"antimalwareDrivers"`
					XenDrivers             bool `json:"xenDrivers"`
				} `json:"monitoringOptions"`
				Actions struct {
					PrimaryAction           int `json:"primaryAction"`
					RemediationAction       int `json:"remediationAction"`
					BackupRemediationAction int `json:"backupRemediationAction"`
				} `json:"actions"`
				ForensicOptions struct {
					OsFailures   bool `json:"osFailures"`
					DriverEvents bool `json:"driverEvents"`
				} `json:"forensicOptions"`
			} `json:"kernelSpaceMemoryIntrospection"`
			Exclusions struct {
				Enabled bool          `json:"enabled"`
				Items   []interface{} `json:"items"`
			} `json:"exclusions"`
			CustomTools struct {
				Enabled bool          `json:"enabled"`
				Tools   []interface{} `json:"tools"`
			} `json:"customTools"`
		} `json:"hvi"`
		StorageProtection struct {
			Icap struct {
				Enabled       bool   `json:"enabled"`
				ServiceName   string `json:"serviceName"`
				EnablePort    bool   `json:"enablePort"`
				ListenPort    int    `json:"listenPort"`
				EnableSslPort bool   `json:"enableSslPort"`
				ListenSslPort int    `json:"listenSslPort"`
				ArchiveScan   struct {
					Enabled         bool `json:"enabled"`
					ArchiveMaxSize  int  `json:"archiveMaxSize"`
					ArchiveMaxDepth int  `json:"archiveMaxDepth"`
				} `json:"archiveScan"`
				CongestionControl struct {
					Mode           int `json:"mode"`
					MaxConnections int `json:"maxConnections"`
				} `json:"congestionControl"`
				ScanAction struct {
					DefaultAction int `json:"defaultAction"`
				} `json:"scanAction"`
			} `json:"icap"`
			Exclusions struct {
				Enabled        bool          `json:"enabled"`
				ExclusionItems []interface{} `json:"exclusionItems"`
			} `json:"exclusions"`
		} `json:"storageProtection"`
		EdrSensor struct {
			General struct {
				Enabled bool `json:"enabled"`
			} `json:"general"`
		} `json:"edrSensor"`
		IntegrityMonitor struct {
			RealTime struct {
				Enabled bool `json:"enabled"`
			} `json:"realTime"`
		} `json:"integrityMonitor"`
	} `json:"uiSettings"`
	Service int `json:"service"`
}
