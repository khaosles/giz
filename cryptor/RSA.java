package cryptor;
/*
 *  File: RsaWithGO.java
 *  Author: khaosles
 *  Date: 2023/6/21 11:26
 *  Desc:
 */

import javax.crypto.BadPaddingException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;
import java.nio.charset.StandardCharsets;
import java.security.*;
import java.security.interfaces.RSAPrivateKey;
import java.security.interfaces.RSAPublicKey;
import java.security.spec.InvalidKeySpecException;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;
import java.util.Base64;


public class RSA {
    /**
     * 与go可以通用的rsa加解密方法
     */
    public static void main(String[] args) throws Exception {
        try {
            String publicKeyString = "MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAKCK8OC6pG8lw+x9YQPeoxclxVJbTJDCKhBHpgCFY5QwlYIZ1aDGXFq0Ui9y4gCJzDkFQwJinq7DC0cze2z+Zi0CAwEAAQ==";
            String privateKeyString = "MIIBVgIBADANBgkqhkiG9w0BAQEFAASCAUAwggE8AgEAAkEAoIrw4LqkbyXD7H1hA96jFyXFUltMkMIqEEemAIVjlDCVghnVoMZcWrRSL3LiAInMOQVDAmKersMLRzN7bP5mLQIDAQABAkB13Wv5Ya3lmDFeh7JZN/Q+w1E6JKNyx8pAB4o9abDxEwuOoTun4lqbEjGbe8MaJlafbmwqCJoeHsZd1lcfBtbVAiEA3DYS/9vbRP/aFHmepZSIdK2c7SDfKs6QGMZ3u8jIkN8CIQC6ollYXMSOk4oCIc7xugYjFFk7g9XNyjLvHZ+B1IrucwIhAIRtsHdVzENuacOucA27iWRBlAG2pv6jQlzj3dV8JcGZAiEAmllZ+rH9/MwX/ttnApMOMETL52SIlcD7jrW5DO5yV/ECIQCQwh7iFuZM3MvwZLRgRbThpfFPygJVWNUuMNjNNYXn1Q==";
            String data = "Hello, JAVA!";
            System.out.println("公钥: " + publicKeyString);
            System.out.println("私钥: " + privateKeyString);

            String encrypted = publicKeyEncrypt(data, publicKeyString);
            System.out.println("java密文: " + encrypted);
            // go 生成的密文
            encrypted = "coX3MkD3ino5Qstq5ZRKetYysnRr3nta96SA7CH4fv4aoQIo5YvzzHFEsE2XZuJ45bEAi8GF9BwGyeT9CfHU9Q==";
            String dataGo = privateKeyDecrypt(encrypted, privateKeyString);
            System.out.println("解密数据： " + dataGo);

            String sign = sign(data, privateKeyString);
            System.out.println("java签名: " + sign);
            // go 生成签名
            sign = "Y0OjWH1WyZ0eQPomaG/T0m3tGvUqY3bFO4Sq2zSBNTu52kkb717mp66M/9b1hgKgaOY0ktE+SjjPDXPb3DihkQ==";
            System.out.println("签名验证结果： " + verify(dataGo, sign, publicKeyString));

        } catch (NoSuchAlgorithmException e) {
            System.out.println(e.getMessage());
        }
    }

    // 生成key
    public static void genKey(int bit) throws NoSuchAlgorithmException {
        KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA");
        keyPairGen.initialize(bit, new SecureRandom());
        KeyPair keyPair = keyPairGen.generateKeyPair();
        RSAPrivateKey privateKey = (RSAPrivateKey) keyPair.getPrivate();
        RSAPublicKey publicKey = (RSAPublicKey) keyPair.getPublic();
        //公钥
        String publicKeyString = Base64.getEncoder().encodeToString(publicKey.getEncoded());
        //私钥
        String privateKeyString = Base64.getEncoder().encodeToString(privateKey.getEncoded());
        System.out.println("公钥: " + publicKeyString);
        System.out.println("私钥: " + privateKeyString);
    }

    /**
     * 加密
     *
     * @param data            数据
     * @param publicKeyString 公钥
     * @return 加密后数据
     */
    public static String publicKeyEncrypt(String data, String publicKeyString) throws NoSuchAlgorithmException, InvalidKeySpecException, IllegalBlockSizeException, BadPaddingException, InvalidKeyException, NoSuchPaddingException {
       // 解码公钥字符串
        byte[] decoded = Base64.getDecoder().decode(publicKeyString);

        // 获取公钥对象
        RSAPublicKey pubKey = (RSAPublicKey) KeyFactory.getInstance("RSA").generatePublic(new X509EncodedKeySpec(decoded));

        // 创建 Cipher 对象并进行加密初始化
        Cipher cipher = Cipher.getInstance("RSA");
        cipher.init(Cipher.ENCRYPT_MODE, pubKey);

        // 执行加密操作并返回加密后的密文
        byte[] encryptedBytes = cipher.doFinal(data.getBytes(StandardCharsets.UTF_8));
        return Base64.getEncoder().encodeToString(encryptedBytes);
    }

    /**
     * 解密
     *
     * @param data             密文
     * @param privateKeyString 私钥
     * @return 解密结果
     */
    public static String privateKeyDecrypt(String data, String privateKeyString) throws NoSuchAlgorithmException, InvalidKeySpecException, NoSuchPaddingException, IllegalBlockSizeException, BadPaddingException, InvalidKeyException {
        // 解码密文和私钥字符串
        byte[] inputByte = Base64.getDecoder().decode(data.getBytes(StandardCharsets.UTF_8));
        byte[] decodedPrivKey = Base64.getDecoder().decode(privateKeyString);

        // 获取私钥对象
        RSAPrivateKey privKey = (RSAPrivateKey) KeyFactory.getInstance("RSA").generatePrivate(new PKCS8EncodedKeySpec(decodedPrivKey));

        // 创建 Cipher 对象并进行解密初始化
        Cipher cipher = Cipher.getInstance("RSA");
        cipher.init(Cipher.DECRYPT_MODE, privKey);

        // 执行解密操作并返回解密后的明文
        byte[] decryptedBytes = cipher.doFinal(inputByte);
        return new String(decryptedBytes, StandardCharsets.UTF_8);
    }

    /**
     * 生成签名
     *
     * @param data             数据
     * @param privateKeyString 私钥
     * @return 签名
     */
    public static String sign(String data, String privateKeyString) throws NoSuchAlgorithmException, InvalidKeySpecException, InvalidKeyException, SignatureException {
        byte[] decodedPrivKey = Base64.getDecoder().decode(privateKeyString);
        RSAPrivateKey privateKey = (RSAPrivateKey) KeyFactory.getInstance("RSA").generatePrivate(new PKCS8EncodedKeySpec(decodedPrivKey));
        // 创建签名对象并初始化
        Signature signature = Signature.getInstance("SHA256withRSA");
        signature.initSign(privateKey);
        // 更新要签名的数据
        signature.update(data.getBytes());
        // 进行签名
        byte[] signatureBytes = signature.sign();
        return Base64.getEncoder().encodeToString(signatureBytes);
    }

    /**
     * 验证签名
     *
     * @param data            数据
     * @param signString      签名
     * @param publicKeyString 公钥
     * @return true 验证成功 false失败
     */
    public static Boolean verify(String data, String signString, String publicKeyString) throws NoSuchAlgorithmException, InvalidKeyException, SignatureException, InvalidKeySpecException {
        byte[] decoded = Base64.getDecoder().decode(publicKeyString);
        RSAPublicKey publicKey = (RSAPublicKey) KeyFactory.getInstance("RSA").generatePublic(new X509EncodedKeySpec(decoded));
        // 验证签名
        Signature signature = Signature.getInstance("SHA256withRSA");
        signature.initVerify(publicKey);
        signature.update(data.getBytes());
        byte[] signatureBytes = Base64.getDecoder().decode(signString.getBytes());
        return signature.verify(signatureBytes);
    }
}
