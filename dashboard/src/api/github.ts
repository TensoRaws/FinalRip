import type { SelectOption } from 'naive-ui'

export async function getGitHubTemplates(
  repo: string,
  token: string = '',
  extension: string = 'py',
): Promise<SelectOption[]> {
  const templateList: SelectOption[] = []
  const url = `https://api.github.com/repos/${repo}/contents/templates`

  const response = await fetch(url, {
    headers: {
      Authorization: token ? `token ${token}` : '',
    },
  })
  const data = await response.json()

  data.forEach((item: any) => {
    if (item.name.split('.').pop() === extension) {
      templateList.push({
        label: item.name,
        value: item.download_url,
      })
    }
  })

  return templateList
}

export async function getGitHubTemplateContent(template: SelectOption): Promise<string> {
  const response = await fetch(String(template?.value))
  return response.text()
}
