# FastHealth-Tools
<p align="center">
    <img
      alt="FastHealth"
      src="fasthealth_logo.svg"
      width="200"
    />
</p>

FastHealth-Tools is a set of open source data munging tools (and utilities) for healthcare data provided by [VertisPro][]. 

The aim is to provide a set of simple but powerful tools that can operate on healthcare data to carry out identification, cleanup, structuring and transformation operations which are often required when preparing data for use in analytics, machine learning and AI.

Some tools may require an understanding of healthcare IT topics like [FHIR][] and terminologies such as [SNOMED CT][], [ICD-10][] and [CPT][].

## Tools

### SSPAF: Super Simple PBKDF2 + AES-128 encryption and decryption for Files
The tool could be used in cases where large amounts of sensitive data needs to be encrypted and stored in seperate files.
* Encrypts and decrypts a file using a passphrase supplied . 
* PBKDF2 is used for generating key from the passphrase.
* AES-128 is used for data encryption.
* There is currently no consensus for an encryted file format so the utility implements its own.
* The utility is meant for automated / batch usage and does not prompt while over writing files.

Command Line Usage:
```Shell
-p  (Required) Passphrase for encrypting or decrypting the file
-i  (Required) Input file
-o  (Required) Output file
-d  Decrypt - by defualt the utility will only encrypt the file
```
Encryption Example:
```Shell
sspaf -p klJ9@0823r2$hk -i patients_otp.csv -o patients_otp_csv.enc
```

Decryption Example:
```Shell
sspaf -d -p klJ9@0823r2$hk -i patients_otp.csv -o patients_otp_csv.enc
```
### (Under Construction) SHNM - Sensible Human Name Masking Tool
The tool takes in human names and genders and produces sensible human names along with masking information.
For example, one could effectively mask patient, next-of-kin, clinican names with sensible sounding names and provide corelated informaion with the mask so that data could be re-identified in the future. This (sort of) ensures compliance with certain GDPR requirements where data can be retroactively deleted from the extracted set at a later time by simply providing a ID that needs to be deleted.

* The correlation is provided using Universally Unique Lexicographically Sortable Identifier ([ULID][])
* Ability to generate more than 5 million unique and sensible names in each gender (capacity can be easily augmented)
* Provides name clash avoidance (requires tool config and state to be saved)

[VertisPro]: https://www.vertispro.com
[FHIR]: https://www.hl7.org/fhir
[SNOMED CT]: https://www.snomed.org/snomed-ct
[ICD-10]: https://en.wikipedia.org/wiki/ICD-10
[CPT]: https://en.wikipedia.org/wiki/Current_Procedural_Terminology
[ULID]: https://github.com/ulid/spec
