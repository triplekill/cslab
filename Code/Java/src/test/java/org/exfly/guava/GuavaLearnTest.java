package org.exfly.guava;

import com.google.common.base.Optional;
import com.google.common.collect.ImmutableSet;
import org.junit.Assert;
import org.junit.Test;


/**
 * @doc https://willnewii.gitbooks.io/google-guava/content
 */
public class GuavaLearnTest {
    @Test
    public void testOptional() {
        Optional possible = Optional.of(5);
        Assert.assertEquals(true, possible.isPresent());
        Assert.assertEquals(5, possible.get());
    }

    @Test
    public void testImmutableCollection() {
        ImmutableSet<String> COLOR_NAMES = ImmutableSet.of(
                "red",
                "orange",
                "purple");
        System.out.println(COLOR_NAMES);
    }
}
