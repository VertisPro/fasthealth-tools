# FastHealth-Tools
<p align="center">
    <img
      alt="FastHealth"
      src="fasthealth_logo.svg"
      width="200"
    />
</p>

FastHealth-Tools is a set of data munging tools and utilities for healthcare data provided by [VertisPro][]. 

These tools provide a preliminary set of data identification and cleaning tools that are often required when dealing with healthcare data in machine learning and data analytics.

In case you wish to get started, it may help to readup on FHIR and terminologies such as SNOMED CT, ICD10 and CPT.

## Utilities

### filecrypt
filecrypt encrypts and decrypts a file using a passphrase supplied to it. 
* PBHDF2 is used for generating key from the passphrase and the file is encrytped using AES-128. 
* There is currently no consensus for an encryted file format so the utility implements its own.
* The utility is meant for automated usage and does not prompt while over writing files.

[VertisPro]: https://vertispro.com
