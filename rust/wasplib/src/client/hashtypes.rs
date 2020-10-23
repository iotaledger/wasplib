#[derive(Eq, PartialEq)]
pub struct ScAddress {
    address: String,
}

impl ScAddress {
    pub fn from_bytes(bytes: &str) -> ScAddress {
        ScAddress { address: bytes.to_string() }
        //ScAddress { address: bytes.try_into().expect("address should be 33 bytes") }
    }

    pub fn to_bytes(&self) -> &str {
        &self.address
    }

    pub fn to_string(&self) -> String {
        self.address.to_string()
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Eq, PartialEq)]
pub struct ScColor {
    color: String,
}

impl ScColor {
    pub fn iota() -> ScColor {
        ScColor { color: "iota".to_string() }
    }

    pub fn mint() -> ScColor {
        ScColor { color: "new".to_string() }
    }

    pub fn from_bytes(bytes: &str) -> ScColor {
        ScColor { color: bytes.to_string() }
        //ScColor { color: bytes.try_into().expect("color should be 32 bytes") }
    }

    pub fn to_bytes(&self) -> &str {
        &self.color
    }

    pub fn to_string(&self) -> String {
        self.color.to_string()
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Eq, PartialEq)]
pub struct ScRequestId {
    id: String,
}

impl ScRequestId {
    pub fn from_bytes(bytes: &str) -> ScRequestId {
        ScRequestId { id: bytes.to_string() }
        //ScRequestId { id: bytes.try_into().expect("request id should be 34 bytes") }
    }

    pub fn to_bytes(&self) -> &str {
        &self.id
    }

    pub fn to_string(&self) -> String {
        self.id.to_string()
    }

    // pub fn transaction_id(&self) -> ScTransactionId {
    //     ScTransactionId::from_bytes(&self.id[..32])
    // }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Eq, PartialEq)]
pub struct ScTransactionId {
    id: String,
}

impl ScTransactionId {
    pub fn from_bytes(bytes: &str) -> ScTransactionId {
        ScTransactionId { id: bytes.to_string() }
        //ScTransactionId { id: bytes.try_into().expect("transaction id should be 32 bytes") }
    }

    pub fn to_bytes(&self) -> &str {
        &self.id
    }

    pub fn to_string(&self) -> String {
        self.id.to_string()
    }
}

