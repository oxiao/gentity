# gentity

模板中内置的md定义为
&ModelInfo{
		Table:           table,
		ImportName:      importName,
		PackageName:     packageName,
		BDName:          DbName,
		ExportModelName: HumpStructName(table.Name),
		TableName:       table.Name,
		ModelName:       table.Name,
		Columns:         table.Columns(),
		TableComment:    table.Comment,
	}