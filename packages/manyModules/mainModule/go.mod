module mainModule

go 1.24.5

replace importedModule/somePackage => ../importModule

require importedModule/somePackage v0.0.0-00010101000000-000000000000
