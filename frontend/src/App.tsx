import { FormEvent, useState } from 'react'
import { calculate, Operation } from './api/calculator'

const operations: Array<{ value: Operation; label: string }> = [
  { value: 'add', label: 'Addition (+)' },
  { value: 'subtract', label: 'Subtraction (−)' },
  { value: 'multiply', label: 'Multiplication (×)' },
  { value: 'divide', label: 'Division (÷)' },
]

function parseOperand(value: string, label: string): number {
  if (value.trim() === '') {
    throw new Error(`Enter a valid ${label} number.`)
  }

  const number = Number(value)
  if (!Number.isFinite(number)) {
    throw new Error(`Enter a valid ${label} number.`)
  }

  return number
}

export function App() {
  const [firstOperand, setFirstOperand] = useState('')
  const [secondOperand, setSecondOperand] = useState('')
  const [operation, setOperation] = useState<Operation>('add')
  const [result, setResult] = useState<number | null>(null)
  const [error, setError] = useState('')
  const [isCalculating, setIsCalculating] = useState(false)

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setError('')
    setResult(null)

    let a: number
    let b: number

    try {
      a = parseOperand(firstOperand, 'first')
      b = parseOperand(secondOperand, 'second')
    } catch (validationError) {
      setError(validationError instanceof Error ? validationError.message : 'Check the entered numbers.')
      return
    }

    setIsCalculating(true)

    try {
      const calculatedResult = await calculate({ operation, a, b })
      setResult(calculatedResult)
    } catch (requestError) {
      setError(requestError instanceof Error ? requestError.message : 'The calculation could not be completed.')
    } finally {
      setIsCalculating(false)
    }
  }

  return (
    <main className="page-shell">
      <section className="calculator-card" aria-labelledby="calculator-title">
        <div className="heading-group">
          <p className="eyebrow">Technical Assessment</p>
          <h1 id="calculator-title">Full-Stack Calculator</h1>
          <p className="subtitle">Enter two numbers, choose an operation, and let the Go API do the math.</p>
        </div>

        <form onSubmit={handleSubmit} noValidate>
          <label htmlFor="first-operand">First number</label>
          <input
            id="first-operand"
            name="firstOperand"
            type="text"
            inputMode="decimal"
            value={firstOperand}
            onChange={(event) => setFirstOperand(event.target.value)}
            placeholder="e.g. 10"
            autoComplete="off"
          />

          <label htmlFor="operation">Operation</label>
          <select
            id="operation"
            name="operation"
            value={operation}
            onChange={(event) => setOperation(event.target.value as Operation)}
          >
            {operations.map((item) => (
              <option key={item.value} value={item.value}>
                {item.label}
              </option>
            ))}
          </select>

          <label htmlFor="second-operand">Second number</label>
          <input
            id="second-operand"
            name="secondOperand"
            type="text"
            inputMode="decimal"
            value={secondOperand}
            onChange={(event) => setSecondOperand(event.target.value)}
            placeholder="e.g. 5"
            autoComplete="off"
          />

          <button type="submit" disabled={isCalculating}>
            {isCalculating ? 'Calculating…' : 'Calculate'}
          </button>
        </form>

        <div className="feedback" aria-live="polite">
          {error && <p className="error-message" role="alert">{error}</p>}
          {result !== null && (
            <p className="result-message">
              Result: <strong>{result}</strong>
            </p>
          )}
        </div>
      </section>
    </main>
  )
}
