import { afterEach, describe, expect, it, vi } from 'vitest'
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { App } from './App'

describe('App', () => {
  afterEach(() => {
    vi.unstubAllGlobals()
  })

  it('submits the selected operation and displays the result', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ result: 42 }),
    })
    vi.stubGlobal('fetch', fetchMock)
    const user = userEvent.setup()

    render(<App />)

    await user.type(screen.getByLabelText('First number'), '6')
    await user.selectOptions(screen.getByLabelText('Operation'), 'multiply')
    await user.type(screen.getByLabelText('Second number'), '7')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(await screen.findByText(/Result:/)).toHaveTextContent('Result: 42')
    expect(fetchMock).toHaveBeenCalledWith('/api/calculate', expect.objectContaining({
      method: 'POST',
      body: JSON.stringify({ operation: 'multiply', a: 6, b: 7 }),
    }))
  })

  it('shows a validation error before calling the API', async () => {
    const fetchMock = vi.fn()
    vi.stubGlobal('fetch', fetchMock)
    const user = userEvent.setup()

    render(<App />)

    await user.type(screen.getByLabelText('Second number'), '5')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(screen.getByRole('alert')).toHaveTextContent('Enter a valid first number.')
    expect(fetchMock).not.toHaveBeenCalled()
  })

  it('displays a backend error', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: false,
      json: async () => ({ error: 'division by zero' }),
    })
    vi.stubGlobal('fetch', fetchMock)
    const user = userEvent.setup()

    render(<App />)

    await user.type(screen.getByLabelText('First number'), '10')
    await user.selectOptions(screen.getByLabelText('Operation'), 'divide')
    await user.type(screen.getByLabelText('Second number'), '0')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(await screen.findByRole('alert')).toHaveTextContent('division by zero')
  })

  it('shows a friendly message when the API cannot be reached', async () => {
    vi.stubGlobal('fetch', vi.fn().mockRejectedValue(new Error('network error')))
    const user = userEvent.setup()

    render(<App />)

    await user.type(screen.getByLabelText('First number'), '3')
    await user.type(screen.getByLabelText('Second number'), '4')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(await screen.findByRole('alert')).toHaveTextContent('Unable to reach the calculator service.')
  })
})
