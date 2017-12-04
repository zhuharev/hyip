# Changelog

## [0.2.0] 2017-11-04
### Added
- Social links configuration
- Name field for Investment plans
- Telegram binding account. Now telegram bot can send user password (generate
  and rewrite new password).
- Profit history

### Changed
- Conf file name mapper now `snake_case`
- Changed conf format to yaml
- Landing page urls now scrol on page, not a link
- Some buttons on landing page route to /login page
- Reworked trader

## [0.1.0] 2017-11-30
## Added
- Support: tickets and admin.
- Redirect from login and reg pages when already signed in. Redirect from dash,
 if signed out

## [0.0.2] 2017-11-29
### Added
- fixer.io for currency conversation

### Fixed
- fix currency panicking (index out of range error)
- fix GetByAmount order and comparing
- fix adv cash bug, when AmountInUSD emty if txn currency == USD

## [0.0.1] 2017-11-28
Initial
