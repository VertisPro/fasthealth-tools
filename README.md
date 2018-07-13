# FastHealth-Tools
<p align="center">
    <img
      alt="FastHealth"
      src="fasthealth_logo.svg"
      width="200"
    />
</p>

FastHealth-Tools are a set of open source data munging tools (and utilities) for healthcare data provided by [VertisPro][]. 

The aim is to provide a set of simple but powerful tools that can operate on healthcare data to carry out identification, cleanup, structuring and transformation operations which are often required when preparing data for use in analytics, machine learning and AI.

Some tools may require an understanding of healthcare IT topics like [FHIR][] and terminologies such as [SNOMED CT][], [ICD-10][] and [CPT][].

## Tools

### Clinical Text Spell Checker (under construction)
The tool carries out spell check over blobs of clinical text and provides suggestions that may be used to clean up unstructured clinical data. This is especially useful in cases where a part of clinical data comes from summaries, notes and reports (mostly in lab and radiology environments) and needs to be cleaned before processing into tools like [Word2vec][], clinical context determination (part of fasthealth). A large dictionary of (US/AU) medical and (US) english words has been provided and it is easy to plug in your own. The input needs is a csv file with column information and the output spelling suggestions are provided as an extra column in the same csv file for easy correlation. This tool uses the ENCHANT library (2.2.1) which needs to be compiled for your platform and linked with the app.

### Sensible Human Name Masking Tool (SHNM)
The tool produces sensible human names along with masking information. This could be used for example to effectively mask patient, next-of-kin, clinican names with sensible sounding names and provide corelated informaion with the mask so that data could be re-identified in the future. This (sort of) ensures compliance with certain GDPR requirements where data can be retroactively deleted from the extracted set at a later time by simply providing a ID that needs to be deleted.

* The correlation is provided using Universally Unique Lexicographically Sortable Identifier ([ULID][])
* Ability to generate ~more than 5~ upto 25  million unique and sensible names in each gender (capacity can be easily augmented).
* Provides name clash avoidance and provides (2) vectors in case patients need to be added later.
* ULID's are unique and lexicographically sortable and can be used as demographic identifiers.

 The ULI
Command Line Usage:
```bash
-n  Number of human names to be generated (max 25 million)
-g Gender of the human names to be generated 'f' for females and 'm' for males. will generate female if unspecified
-u Generate ULID's (false by default), set to true if you want them
```
Example generating 100 unique male names with ULID's :
```bash
shnm -n 100 -g m -u
```

### Simple PBKDF2 + AES-128 encryption and decryption for Files (SPAF)
The tool could be used in cases where large amounts of sensitive data needs to be encrypted and stored in seperate files.
* Encrypts and decrypts a file using a passphrase supplied . 
* PBKDF2 is used for generating key from the passphrase.
* AES-128 is used for data encryption.
* There is currently no consensus for an encryted file format so the utility implements its own.
* The utility is meant for automated / batch usage and does not prompt while over writing files.

Command Line Usage:
```bash
-p  (Required) Passphrase for encrypting or decrypting the file
-i  (Required) Input file
-o  (Required) Output file
-d  Decrypt - by defualt the utility will only encrypt the file
```
Encryption Example:
```bash
sspaf -p klJ9@0823r2$hk -i patients_otp.csv -o patients_otp_csv.enc
```

Decryption Example:
```bash
sspaf -d -p klJ9@0823r2$hk -i patients_otp.csv -o patients_otp_csv.enc
```

[VertisPro]: https://www.vertispro.com
[FHIR]: https://www.hl7.org/fhir
[SNOMED CT]: https://www.snomed.org/snomed-ct
[ICD-10]: https://en.wikipedia.org/wiki/ICD-10
[CPT]: https://en.wikipedia.org/wiki/Current_Procedural_Terminology
[ULID]: https://github.com/ulid/spec
[Word2vec]: https://en.wikipedia.org/wiki/Word2vec