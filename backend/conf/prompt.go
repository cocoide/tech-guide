package conf

const SelectTopicsPrompt = "[Title: %s], Select tags related to the title from [%s] with weights ranging from 1 to 5. Output only the tags with a weight of 3 or higher in the format `<TagName>: <Weight>,`(do not output anything for tags with weights below 3)"
