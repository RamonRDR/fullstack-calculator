export type Operation = 'add' | 'subtract' | 'multiply' | 'divide'

export interface CalculationRequest {
  operation: Operation
  a: number
  b: number
}

interface CalculationResponse {
  result: number
}

interface ErrorResponse {
  error?: string
}

export async function calculate(request: CalculationRequest): Promise<number> {
  let response: Response

  try {
    response = await fetch('/api/calculate', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(request),
    })
  } catch {
    throw new Error('Unable to reach the calculator service.')
  }

  const payload = (await response.json().catch(() => ({}))) as CalculationResponse & ErrorResponse

  if (!response.ok) {
    throw new Error(payload.error || 'The calculation could not be completed.')
  }

  if (typeof payload.result !== 'number') {
    throw new Error('The calculator service returned an invalid response.')
  }

  return payload.result
}
