package org.exfly.lombok;

import org.junit.Assert;
import org.junit.Test;

public class LombokLApplicationTest {

    public void contextLoads() {
    }

    @Test
    public void testt() {
        Assert.assertEquals(1, 1);
    }

    @Test
    public void testProduct() {
        Product p = new Product(10, "xxx");
        Assert.assertEquals(10, p.getIa());
    }
}
