context('Basic', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  it('basic nav', () => {
    cy.url()
      .should('eq', 'http://localhost:3333/')

    cy.contains('[Home Layout]')
      .should('exist')

    cy.get('#input').should('be.visible').click()
      .type('ViteTail')

    cy.get('#go').click()
      .url().should('eq', 'http://localhost:3333/hi/ViteTail')

    cy.contains('[Default Layout]')
      .should('exist')
  })
})
